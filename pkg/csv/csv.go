package csv

import (
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"github.com/kierachell/cde-cnn/pkg/data/types"
	"io"
	"log"
	"os"
)

func ReadDataElements(filePath string, limit int, offset int) []types.DataElement {
	csvFile, failedToOpenFile := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if failedToOpenFile != nil {
		log.Printf("Failed to open file: %s", failedToOpenFile)
	}
	defer csvFile.Close()

	var dataElements []types.DataElement
	reader := csv.NewReader(csvFile)
	dec, failedToDecode := csvutil.NewDecoder(reader)

	if failedToDecode != nil {
		log.Printf("Failed to decode: %v", failedToDecode.Error())
	}

	for {
		if offset >= 0 {
			offset--
			continue
		}
		var u types.DataElement
		if readError := dec.Decode(&u); readError == io.EOF {
			break
		} else if readError != nil {
			log.Printf("Failed to read lines: %v", readError.Error())
		}

		dataElements = append(dataElements, u)

		limit--
		if limit == 0 {
			break
		}
	}
	return dataElements
}

func WriteDataElements(filePath string, dataElements []types.DataElement) error {
	file, err := os.OpenFile(filePath+".csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return err
	}

	writer := csv.NewWriter(file)
	enc := csvutil.NewEncoder(writer)

	for _, data := range dataElements {
		if err := enc.Encode(data); err != nil {
			log.Printf("Error encoding data: %v", err)
			return err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Printf("Error writing data: %v", err)
		return err

	}
	return nil
}
