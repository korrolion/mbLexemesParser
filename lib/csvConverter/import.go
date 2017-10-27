package csvConverter

import (
	"io/ioutil"
	"log"
	"encoding/csv"
	"strings"
	"io"
)

func Import() map[string]map[string]string {
	var data, err = ioutil.ReadFile("import/greetup_input.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(data)))

	result := make(map[string]map[string]string)
	langs := []string{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(langs) == 0 {
			//Имеем дело с заголовком
			langs = record
			continue
		}

		var key string
		for i, value := range record {
			if i == 0 {
				key = value
				continue
			}
			lang := langs[i]
			if result[key] == nil {
				result[key] = make(map[string]string)
			}
			result[key][lang] = value
		}

	}

	return result
}