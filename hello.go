package main

import "fmt"
func initValue(){
    var a int
    var s string
    // fmt.Printf("%d %q/n", a, s)
    fmt.Println(a, s)
}

func setValue(){
    var a, b int = 2, 3
    var s string = "abc"
    fmt.Println(a, b, s)
}

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("测试代码")

    initValue()
    setValue()
}
