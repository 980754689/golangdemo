package main

import "fmt"

var aa = 1
var bb = 2

var (
    cc = 3
    dd = "abc"
)

func initValue(){
    var a int
    var s string
    fmt.Printf("%d %q\n", a, s)
    // fmt.Println(a, s)
}

func setValue(){
    var a, b int = 2, 3
    var s string = "abc"
    fmt.Println(a, b, s, dd)
}

func setValueOne(){
    a, b, c, s := 1, 2, false, "abc"
    b = 3
    c = true
    fmt.Println(a, b, c, s, cc)  
}

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("测试代码")

    initValue()
    setValue()
    setValueOne()
    fmt.Println(aa, bb)
}
