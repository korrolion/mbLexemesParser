package main

import (
	"log"
	"fmt"
	"github.com/korrolion/mbParser/lib/csvConverter"
	"regexp"
	"io/ioutil"
)

func exportData() {
	greetupDic := make(lexemsMatrix)

	for _, lang := range langs {
		if contains(newLangs, lang) { continue }
		filename := "source/targetProj/Locale/" + lang + ".lproj/Localizable.strings"
		greetupDic = exportFile(greetupDic, filename, lang)
	}

	greetupDic = match(greetupDic)

	csvConverter.Export(greetupDic, langs)
}

func exportFile(dict lexemsMatrix, filename, lang string) lexemsMatrix {

	var data, err = ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	sourceString := cleanData(data)

	var state IteartionState = waitingForKey
	syntaxChecked := true
	key := ""
	value := ""
	prevCh := ""
	for i, r := range sourceString {
		ch := fmt.Sprintf("%c", r)
		if i >= 1 {
			prevCh = fmt.Sprintf("%c", sourceString[i-1])
		}

		//Валидация синтаксиса
		if ch == ";" && state == waitingForKey {
			syntaxChecked = true
		} else if ch == "=" && state == waitingForValue {
			syntaxChecked = true
		}


		//Переключение стейта
		if ch == "\"" && prevCh != "\\" {

			switch state {
			case waitingForKey:
				if !syntaxChecked {
					log.Fatal("Broken syntax: expected ;")
				}
				state = onKey
			case waitingForValue:
				if !syntaxChecked {
					log.Fatal("Broken syntax: expected =")
				}
				state = onValue
			case onKey:
				state = waitingForValue
				syntaxChecked = false
			case onValue:
				state = waitingForKey

				if dict[key] == nil {
					dict[key] = make(map[string]string)
				}

				dict[key][lang] = value

				key = ""
				value = ""
				syntaxChecked = false
			default:
				log.Fatal("Undefinded Iteration State")
			}
			continue

		}

		//Сбор значений
		switch state {
		case onKey:
			key += ch
		case onValue:
			value += ch

		}
	}

	return dict
}

func cleanData(data []byte) string {
	//Удалим комменты
	re := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	newBytes := re.ReplaceAll(data, nil)
	return string(newBytes)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}