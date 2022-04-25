package main

import (
	"fmt"
	"math/cmplx"
)

func main(){
	c := 3 + 4i
	var b =cmplx.Abs(c)
	fmt.Println(b)
}