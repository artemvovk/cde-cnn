package preprocess

import (
	"log"
	"math"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"unicode"

	"gopkg.in/jdkato/prose.v2"

	"github.com/kierachell/cde-cnn/pkg/data/types"
)

func GetTokens(dataElements []types.DataElement, fieldName string) map[types.DataElement][]types.Token {
	var elementTokenMap sync.Map
	var processed int

	done := make(chan bool)
	defer close(done)
	dataChannel := makePipeline(dataElements, done)

	workers := make([]<-chan int, 8)
	for i := 0; i < len(workers); i++ {
		workers[i] = tokenizeElements(dataChannel, fieldName, done, &elementTokenMap)
	}

	for worked := range fanIn(done, workers...) {
		processed += worked
	}
	log.Printf("Total processed: %v\n", processed)
	returnMap := make(map[types.DataElement][]types.Token, len(dataElements))
	elementTokenMap.Range(func(key, val interface{}) bool {
		returnMap[key.(types.DataElement)] = val.([]types.Token)
		return true
	})

	return returnMap
}

func GetNgrams(sentence string, size int) (count map[string]uint32) {
	reStopwords := MakeStopWords()
	words := splitOnNonLetters(StripStopWords(sentence, reStopwords))
	count = make(map[string]uint32, 0)
	offset := int(math.Floor(float64(size / 2)))

	max := len(words)
	for i, _ := range words {
		if i < offset || i+size-offset > max {
			continue
		}
		gram := strings.Join(words[i-offset:i+size-offset], " ")
		count[gram]++
	}
	return count
}

func MakeStopWords() *regexp.Regexp {
	reStr := `(?i)`
	for _, word := range stopwords {
		reStr += `\b` + word + `\b`
		reStr += `|`
	}
	reStr += `[0-9]+|[,.\-)(]+`

	return regexp.MustCompile(reStr)
}

func StripStopWords(input string, reStopwords *regexp.Regexp) string {
	reInsideWhitespace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)

	input = reStopwords.ReplaceAllString(" "+input, " ")
	input = reInsideWhitespace.ReplaceAllString(input, " ")
	input = strings.TrimSpace(input)

	return input
}

func GetFormNameFromTitle(input string) (string, string) {
	formName := ""
	reSplitOnHypen := regexp.MustCompile(`(?i)(?P<form>.*(\(.*\))?(-| -))(?P<title>.*)`)
	matches := reSplitOnHypen.FindStringSubmatch(input)
	if len(matches) == 0 {
		return input, formName
	}

	for i, match := range reSplitOnHypen.SubexpNames() {
		if match == "form" {
			formName = matches[i]
		}
		if match == "title" {
			input = matches[i]
		}
	}
	return input, formName
}

func makePipeline(dataElements []types.DataElement, done <-chan bool) chan types.DataElement {
	out := make(chan types.DataElement, len(dataElements))
	go func() {
		for _, dataElement := range dataElements {
			select {
			case <-done:
				return
			case out <- dataElement:
			}
		}
		close(out)
	}()
	return out
}

func tokenizeElements(dataElements <-chan types.DataElement, fieldName string, done <-chan bool, elementTokenMap *sync.Map) <-chan int {
	processed := make(chan int)
	var index int
	reStopwords := MakeStopWords()
	go func() {
		for dataElement := range dataElements {
			sentence := getField(&dataElement, fieldName)
			index += 1
			select {
			case <-done:
				return
			case processed <- index:
				sentence = StripStopWords(strings.ToLower(sentence), reStopwords)
				var tokens []types.Token
				doc, _ := prose.NewDocument(sentence)
				for _, token := range doc.Tokens() {
					tokens = append(tokens, types.Token{token.Text, token.Tag, token.Label})
				}
				val, _ := elementTokenMap.Load(dataElement)
				if val != nil {
					elementTokenMap.Store(dataElement, append(val.([]types.Token), tokens...))
				} else {
					elementTokenMap.Store(dataElement, tokens)
				}
			}
		}
		close(processed)
	}()
	return processed
}

func fanIn(done <-chan bool, workers ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(workers))
	result := make(chan int)
	collectWork := func(worker <-chan int) {
		defer wg.Done()
		for work := range worker {
			select {
			case <-done:
				return
			case result <- work:

			}
		}
	}
	for _, worker := range workers {
		go collectWork(worker)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}

func getField(v *types.DataElement, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

func splitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return unicode.IsSpace(char) }
	return strings.FieldsFunc(s, notALetter)
}

var stopwords = [...]string{
	`a`,
	`an`,
	`as`,
	`at`,
	`and`,
	`by`,
	`for`,
	`from`,
	`if`,
	`in`,
	`is`,
	`it`,
	`part`,
	`of`,
	`on`,
	`or`,
	`raw`,
	`score`,
	`to`,
	`the`,
	`that`,
	`with`,
	`score`,
	`scale`,
	`imaging`,
	`indicator`,
	`'s`,
	`.`,
	`,`,
}
