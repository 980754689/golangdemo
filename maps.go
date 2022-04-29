package main

import "fmt"

func main() {
	m := map[string]string{
		"姓名": "张三",
		"性别": "男",
		"年龄": "20",
	}
	fmt.Println(m)

	m2 := make(map[int]string)
	var m3 map[string]int
	fmt.Println(m2, m3)

	//循环m
	for k, v := range m {
		fmt.Println(k,v)
	}

	//取map值
	courseName, ok := m["姓名"]
	fmt.Println(courseName, ok)
	courseSex1, ok := m["性别1"]
	fmt.Println(courseSex1, ok)
	
	if courseSex1, ok := m["性别1"] ; ok {
		fmt.Println(courseSex1)
	} else {
		fmt.Println("key does not exist")
	}

	name ,ok := m["性别"]
	fmt.Println(name,ok)

	delete(m, "性别")
	name ,ok = m["性别"]
	fmt.Println(m, name, ok)








}