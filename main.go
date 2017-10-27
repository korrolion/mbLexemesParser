package main

import (
	"fmt"
	"github.com/korrolion/mbParser/lib/csvConverter"
	"os"
)

type lexemsMatrix map[string]map[string]string

type IteartionState int
const (
	waitingForKey IteartionState = iota
	waitingForValue
	onKey
	onValue
)

func main() {

	args := os.Args[1:]

	if len(args) >= 1 {
		arg := args[0]
		if arg == "import" {
			importData()
		} else if arg == "export" {
			exportData()
		}

		return
	}


	fmt.Println("Need argument export or import")
}

func importData() {
	greetupDic := csvConverter.Import()

	results := make(map[string]string)
	for key, matrix := range greetupDic {
		for lang, value := range matrix {
			results[lang] += "\"" + key + "\"=\"" + value + "\";\n"
		}
	}

	for lang, result := range results {
		saveToFile(lang, result)
	}
}

func saveToFile(lang, result string) {
	// open output file
	fo, err := os.Create("export/targetProj/Locale/" + lang + ".lproj/Localizable.strings")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	//fmt.Fprint(fo, result)
}
