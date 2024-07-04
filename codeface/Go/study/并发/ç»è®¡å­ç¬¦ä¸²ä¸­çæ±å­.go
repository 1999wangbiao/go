package main

import (
	"fmt"
	"unicode"
)

func ChineseCount( s string) (cunt int) {
	for _,char:= range s {
		if unicode.Is(unicode.Han,char) {
			cunt++
		}
	}
	return
}

func main()  {
	fmt.S()
}