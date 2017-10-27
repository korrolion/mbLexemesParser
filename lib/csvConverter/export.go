package csvConverter

import (
	"os"
	"fmt"
        "encoding/csv"
	"log"
	"bytes"
)

func Export(dict map[string]map[string]string, langs []string) {
	records := [][]string{}
	header := []string{""}
	for _, lang := range langs {
		header = append(header, lang)
	}
	records = append(records, header)

	for key, matrix := range dict {
		line := []string{key}
		for _, lang := range langs {
			line = append(line, matrix[lang])
		}
		records = append(records, line)
	}


	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}


	saveToFile(buf.String())
}

func saveToFile(result string) {
	// open output file
	fo, err := os.Create("export/greetup_output.csv")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	fmt.Fprint(fo, result)
}