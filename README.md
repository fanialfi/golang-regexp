# regexp

Go mengadopsi standard **RE2**, untuk melihat syntax yang di support engine ini bisa langsung merujuk ke [dokumentasinya][re2 doc].

## kompilasi ekspresi regexp

function `regexp.Compile(expr string)` digunakan untuk mengkompilasi ekspresi regexp, function tersebut mengembalikan object bertipe `regexp.*Regexp`.

```go
package main

import (
  "fmt"
  "regexp"
)

func main(){
  text := "0Banana 0Curger Soup"
	regexp, err := regexp.Compile(`^\d[a-zA-Z]\w`)

	if err != nil {
		fmt.Println(err.Error())
	}

	res1 := regexp.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)

	res2 := regexp.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
}
```

## function `*regexp.FindAllString(s string, n int)`

digunakan untuk mencari semua string yang sesuai dengan ekspresi regexp, dengan kembalian berupa slice string. 
Jumplah dari hasil pencariannya bisa ditentukan oleh parameter kedua, jika value dari n kurang dari 0, maka akan mengembalikan semua data.

## function `*regexp.MatchString(s string)`

digunakan untuk mendeteksi apakah string memenuhi pola regexp, return value dari method ini berupa nilai boolean.

```go
package main

import (
  "fmt"
  "regexp"
)

func main(){
  text := "Banana bUrgeR souP"
	regexp, err := regexp.Compile(`[a-z]\w+`)

	if err != nil {
		panic(fmt.Sprintf("panic error occured, where message is %s", err.Error()))
	}

	isMatch := regexp.MatchString(text)
	fmt.Println(isMatch)
}
```

## function `*regexp.FindString(s string)`

sama seperti function `*regexp.FindAllString(s string, n int)`, namun function ini hanya mengembalikan 1 buah saja. 
Jika ada banyak string yang sesuai dengan ekspresi regexp, maka yang dikembalikan hanya yang pertama saja.

## function `*regexp.ReplaceAllString(src, repl string)`

digunakan untuk mereplace semua string yang memenuhi kriterian regexp dengan string lain.

```go
package main

import (
  "fmt"
  "regexp"
)

func main(){
  text := "Banana Melon ApplE"
	regexp, err := regexp.Compile(`[a-z]\w+`)

	if err != nil {
		panic(fmt.Sprintf("panic error occured with message %s\n", err.Error()))
	}

	str := regexp.ReplaceAllString(text, "Potato")
	fmt.Println(str)
}
```

## function `*regexp.ReplaceAllStringFunc(src string, repl func(string) string)`

method ini sama seperti `*regexp.ReplaceAllString()` namun kondisinya bisa ditentukan sendiri setiap sub-string yang akan direplace.

```go
package main

import (
	"fmt"
	"regexp"
)

func main(){
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
```

[re2 doc]: https://github.com/google/re2/wiki/Syntax