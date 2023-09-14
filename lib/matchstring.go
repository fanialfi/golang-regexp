package lib

import (
	"fmt"
	"regexp"
)

func MatchString() {
	text := "Banana bUrgeR souP"
	regexp, err := regexp.Compile(`[a-z]\w+`)

	if err != nil {
		panic(fmt.Sprintf("panic error occured, where message is %s", err.Error()))
	}

	isMatch := regexp.MatchString(text)
	fmt.Println(isMatch)
}
