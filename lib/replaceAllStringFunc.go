package lib

import (
	"fmt"
	"regexp"
)

func ReplaceAllStringFunc() {
	text := "Banana Melon Apple"

	regexp, err := regexp.Compile(`[a-zA-Z]\w+`)

	if err != nil {
		panic(fmt.Sprintf("panic occured where message is %s\n", err))
	}

	str := regexp.ReplaceAllStringFunc(text, func(each string) string {
		switch each {
		case "Banana":
			return "Pisang"
		case "Melon":
			return "Melon"
		case "Apple":
			return "Apel"
		default:
			return each
		}
	})

	fmt.Println(str)
}
