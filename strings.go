package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello word 你好世界"
	fmt.Println(len(s))
	fmt.Printf("%d ", []byte(s))
	fmt.Printf("%s ", []byte(s))
	fmt.Printf("%X ", []byte(s))
	fmt.Println("")
	for _, v := range []byte(s) {
		fmt.Printf("%X ", v)
	}
	fmt.Println("")
	for i, val := range []byte(s) {
		fmt.Printf("(%d %X) ", i, val)
	}
	fmt.Println("")
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	
	for len(bytes)>0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println("")
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	
}