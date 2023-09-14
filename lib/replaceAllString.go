package lib

import (
	"fmt"
	"regexp"
)

func ReplaceAllString() {
	text := "0Banana 0Melon ApplE"
	regexp, err := regexp.Compile(`[a-z]\w+`)

	if err != nil {
		panic(fmt.Sprintf("panic error occured with message %s\n", err.Error()))
	}

	str := regexp.ReplaceAllString(text, "Potato")
	fmt.Println(str)
}
