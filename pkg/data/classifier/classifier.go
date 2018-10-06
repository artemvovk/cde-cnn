package classifier

import (
	"strings"

	"github.com/xrash/smetrics"

	"github.com/kierachell/cde-cnn/pkg/data/preprocess"
	"github.com/kierachell/cde-cnn/pkg/data/types"
)

type CDEMatch struct {
	CDE         string
	DataElement types.DataElement
}

func TermFrequencyByDocument(tokensByElement map[types.DataElement][]types.Token) map[types.DataElement]map[string]types.TermFrequency {
	tokenFrequenciesByElement := make(map[types.DataElement]map[string]types.TermFrequency, len(tokensByElement))

	for element, tokens := range tokensByElement {
		for _, token := range tokens {
			if tokenFrequencies, infoExists := tokenFrequenciesByElement[element]; infoExists {
				if word, wordExists := tokenFrequencies[token.Text]; wordExists {
					word.Frequency += 1
				} else {
					tokenFrequencies[token.Text] = types.TermFrequency{
						Word:      token,
						Frequency: 1,
					}
				}
			} else {
				tokenFrequenciesByElement[element] = make(map[string]types.TermFrequency, len(tokens))
				tokenFrequenciesByElement[element][token.Text] = types.TermFrequency{
					Word:      token,
					Frequency: 1,
				}
			}
		}
	}
	return tokenFrequenciesByElement
}

func DocumentFrequencyByTerm(tokensByElement map[types.DataElement][]types.Token) map[types.Token]map[string]types.DocFrequency {
	docFrequenciesByToken := make(map[types.Token]map[string]types.DocFrequency, len(tokensByElement))

	for element, tokens := range tokensByElement {
		for _, token := range tokens {
			if docFrequencies, infoExists := docFrequenciesByToken[token]; infoExists {
				if doc, docExists := docFrequencies[element.Definition]; docExists {
					doc.Frequency += 1
				} else {
					docFrequencies[element.Definition] = types.DocFrequency{
						Doc:       element,
						Frequency: 1,
					}
				}
			} else {
				docFrequenciesByToken[token] = make(map[string]types.DocFrequency, len(tokens))
				docFrequenciesByToken[token][element.Definition] = types.DocFrequency{
					Doc:       element,
					Frequency: 1,
				}
			}
		}
	}
	return docFrequenciesByToken
}

func MatchByTitle(known, data []types.DataElement, minMatch, maxMatch float64) map[string][]types.DataElement {
	sampleSize := len(data)
	dataElementMatches := make(map[string][]types.DataElement)

	cdes := make(chan types.DataElement, sampleSize)
	processLog := make(chan CDEMatch, sampleSize)

	go func() {
		for _, cde := range known {
			cdes <- cde
		}
		close(cdes)
	}()

	go func() {
		reStopwords := preprocess.MakeStopWords()
		for cde := range cdes {
			for j, _ := range data {
				title1, formName1 := preprocess.GetFormNameFromTitle(cde.Title)
				title2, formName2 := preprocess.GetFormNameFromTitle(data[j].Title)
				titleDiff := smetrics.JaroWinkler(preprocess.StripStopWords(strings.ToLower(title1), reStopwords), preprocess.StripStopWords(strings.ToLower(title2), reStopwords), 0.7, 4)
				formNameDiff := smetrics.JaroWinkler(formName1, formName2, 0.7, 4)
				var match CDEMatch
				if titleDiff > minMatch && titleDiff < maxMatch && (formNameDiff > minMatch || formName1 == formName2) {
					match.CDE = cde.VariableName
					match.DataElement = data[j]
					processLog <- match
				}
			}
		}
		defer close(processLog)
	}()

	for match := range processLog {
		dataElementMatches[match.CDE] = append(dataElementMatches[match.CDE], match.DataElement)
	}

	return dataElementMatches
}
