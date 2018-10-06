package main

import (
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/kierachell/cde-cnn/pkg/csv"
	"github.com/kierachell/cde-cnn/pkg/data/classifier"
	"github.com/kierachell/cde-cnn/pkg/data/preprocess"
	"github.com/olekukonko/tablewriter"
)

func TestCsvReader(t *testing.T) {
	dataElements := csv.ReadDataElements("data/tbi_code_cdes.csv", -1, 0)
	log.Printf("%q\t", dataElements[0].Title)
	log.Printf("%q\n", dataElements[0].Definition)
}

func BenchmarkDataElementPreprocess(b *testing.B) {
	sampleSize := 5
	dataElements := csv.ReadDataElements("data/de_export.csv", sampleSize, 0)

	for n := 0; n < b.N; n++ {
		entities := preprocess.GetTokens(dataElements, "Definition")
		for _, ent := range entities {
			log.Printf("%q\n", ent)
		}
	}
}

func TestNgrams(t *testing.T) {
	sampleSize := 200
	gramSize := math.Log10(float64(sampleSize))
	dataElements := csv.ReadDataElements("data/de_export.csv", sampleSize, 0)
	allGrams := make(map[string]uint32, 0)
	for _, dataElement := range dataElements {
		grams := preprocess.GetNgrams(dataElement.Definition, int(gramSize))
		for gram, count := range grams {
			allGrams[gram] += count
		}
	}
	for gram, count := range allGrams {
		if count > uint32(gramSize) {
			log.Printf("%v - %v\n", gram, count)
		}
	}
}

func TestTitleDiff(t *testing.T) {
	sampleSize := 100
	rand.Seed(time.Now().Unix())
	offset := rand.Intn(sampleSize * 2)
	offset = 0

	log.Printf("Offset: %v", offset)
	cdes := csv.ReadDataElements("data/tbi_code_cdes.csv", -1, 0)
	dataElements := csv.ReadDataElements("data/de_export.csv", sampleSize, offset)
	cdeDataElements := classifier.MatchByTitle(cdes, dataElements, 0.7, 0.97)

	for cde, dataElement := range cdeDataElements {
		log.Printf("CDE got mapped: %v", cde)
		csv.WriteDataElements("out/"+cde, dataElement)
	}
}

func TestTermFrequency(t *testing.T) {
	sampleSize := 100
	dataElements := csv.ReadDataElements("data/de_export.csv", sampleSize, 0)
	tokens := preprocess.GetTokens(dataElements, "Definition")
	tokenFrequencies := classifier.TermFrequencyByDocument(tokens)
	for element, tokens := range tokenFrequencies {
		log.Printf("Element Definition: %v", element.Definition)
		for _, info := range tokens {
			log.Printf("%q appears %v times\n", info.Word, info.Frequency)
		}
	}
}

func TestDocFrequency(t *testing.T) {
	sampleSize := 2000
	dataElements := csv.ReadDataElements("data/de_export.csv", sampleSize, 0)
	tokens := preprocess.GetTokens(dataElements, "Definition")
	docFrequencies := classifier.DocumentFrequencyByTerm(tokens)

	tabw := tablewriter.NewWriter(os.Stdout)
	tabw.SetHeader([]string{"Text", "Label", "Document Count"})

	for token, elements := range docFrequencies {
		if len(elements) > 2 {
			tabw.Append([]string{token.Text, token.Tag, strconv.Itoa(len(elements))})
		}
	}
	tabw.Render()
}
