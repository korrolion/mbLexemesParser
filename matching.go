package main

import "fmt"

func match(greetupDic lexemsMatrix) lexemsMatrix {

	mambaDic := make(lexemsMatrix)

	for _, lang := range langs {
		filename := "source/sourceProj/Locale/" + lang + ".lproj/Localizable.strings"
		mambaDic = exportFile(mambaDic, filename, lang)
	}

	matchMap := makeMatchMap("ru", greetupDic, mambaDic)

	for key, matrix := range mambaDic {
		if _, isset := matchMap[key]; !isset { continue }

		greetupKey := matchMap[key]

		for lang, value := range matrix {
			if _, isset := greetupDic[greetupKey][lang]; !isset {
				fmt.Println("Match key", greetupKey, "value", value, "lang", lang)
				greetupDic[greetupKey][lang] = value
			}
		}
	}

	return greetupDic
}


//Словарь матчинга ключ в мамбе ->  ключ в гритапе на одинаковые лексемы
func makeMatchMap(lang string, greetupDic, mambaDic lexemsMatrix) map[string]string {

	matchMap := make(map[string]string)

	for mambaKey, mambaMatrix := range mambaDic {
		for greetupKey, greetupMatrix := range greetupDic {
			if greetupMatrix[lang] == mambaMatrix[lang] {
				//Нашли матч
				//fmt.Println("Match key", greetupKey, "value", mambaMatrix[lang])
				matchMap[mambaKey] = greetupKey
			}
		}

	}

	return matchMap
}