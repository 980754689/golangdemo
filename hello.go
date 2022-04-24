package main

import "fmt"
func initValue(){
    var a int
    var s string
    fmt.Printf("%d %q\n", a, s)
    // fmt.Println(a, s)
}

func setValue(){
    var a, b int = 2, 3
    var s string = "abc"
    fmt.Println(a, b, s)
}

func setValueOne(){
    a, b, c, s := 1, 2, false, "abc"
    b = 3
    c = true
    a = 22
    s = 33
    fmt.Println(a, b, c, s)  
}

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("测试代码")

    initValue()
    setValue()
    setValueOne()
}
