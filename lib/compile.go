package lib

import (
	"fmt"
	"regexp"
)

func Compile() {
	text := "0Banana 0Curger Soup"
	regexp, err := regexp.Compile(`^\d[a-z|A-Z]\w+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	res1 := regexp.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)

	res2 := regexp.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
}
