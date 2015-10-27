package main

import (
	"fmt"
	"unicode/utf8"
)

const (
	NEUTER    = iota // средний
	MASCULINE        // мужской
	FEMININE         // женский
)

func main() {
	words := []string{"горячая", "вода", "холодная", "квартплата", "вывоз", "мусор", "электроэнергия", "отопление"}
	for _, word := range words {

		fmt.Print(word, ": ")
		switch detectKindOfWord(word) {
		case NEUTER:
			fmt.Print("средний")
		case MASCULINE:
			fmt.Print("мужской")
		case FEMININE:
			fmt.Print("женский")
		}
		fmt.Println()
		fmt.Println(detectDeclension(word))
		fmt.Println(ToAccusative(word))
		fmt.Println()
	}
}

func detectKindOfWord(word string) uint8 {
	_, last_rune_size := utf8.DecodeLastRuneInString(word)
	//	_, last_rune_size2 := utf8.DecodeLastRuneInString(word[0 : len(word)-last_rune_size])
	//	suffix2 := word[len(word)-(last_rune_size+last_rune_size2):]
	//	prefix2 := word[0 : len(word)-(last_rune_size+last_rune_size2)]

	suffix1 := word[len(word)-last_rune_size:]
	//	prefix1 := word[0 : len(word)-last_rune_size]

	switch suffix1 {
	case "о", "е":
		return NEUTER
	case "а", "я":
		return FEMININE
	case "й", "ц", "к", "н", "г", "ш", "щ", "з", "х", "ъ", "ф", "в", "п", "р", "л", "д", "ж", "ч", "с", "м", "т":
		return MASCULINE
	default:
		return 10
	}

	return NEUTER
}

func detectDeclension(word string) uint8 {
	_, last_rune_size := utf8.DecodeLastRuneInString(word)
	suffix1 := word[len(word)-last_rune_size:]

	wordKind := detectKindOfWord(word)

	switch {
	case suffix1 == "а" || suffix1 == "я":
		return 1
	case wordKind == MASCULINE || wordKind == NEUTER:
		return 2
	case wordKind == FEMININE:
		return 3
	}
	return 0
}

func ToAccusative(word string) string {
	_, last_rune_size := utf8.DecodeLastRuneInString(word)
	_, last_rune_size2 := utf8.DecodeLastRuneInString(word[0 : len(word)-last_rune_size])
	suffix2 := word[len(word)-(last_rune_size+last_rune_size2):]
	prefix2 := word[0 : len(word)-(last_rune_size+last_rune_size2)]

	// для прилагательных
	switch suffix2 {
	case "ая": // женский род
		return prefix2 + "ую"
	case "яя": // женский род
		return prefix2 + "юю"
	case "ое": // средний род
		return prefix2 + "ое"
	case "ее": // средний род
		return prefix2 + "ее"
	case "ой": // мужской род
		return prefix2 + "ого" // needs check
	case "ый": // мужской род
		return prefix2 + "ого" // needs check
	case "ий": // мужской род
		return prefix2 + "его" // needs check
	}

	suffix1 := word[len(word)-last_rune_size:]
	prefix1 := word[0 : len(word)-last_rune_size]

	declension := detectDeclension(word)
	switch {
	case suffix1 == "а" && declension == 1:
		return prefix1 + "у"
	case suffix1 == "я" && declension == 1:
		return prefix1 + "ю"
	case declension == 2 && suffix1 != "о" && suffix1 != "е":
		return word + "а"
	case declension == 2 && suffix1 == "о":
		return prefix1 + ""
	case declension == 2 && suffix1 == "е":
		return prefix1 + "я"
	}

	return word
}
