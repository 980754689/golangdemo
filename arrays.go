package main

import "fmt"

/*
 * 数组操作
 */

 func printArray(arr *[5]int){
	 arr[0] = 100
	for i, v := range arr {
		fmt.Println(i,v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{3,4,5,6,7}
	var grid [4][5]int

    fmt.Println(arr1, arr2, arr3, grid)
	for i, _ := range arr3 {
		fmt.Println(i)
	}

	//取数组中最大值
	maxi := -1
	maxValue := -1
	for i,v := range arr3 {
		if v > maxValue  {
			maxi , maxValue = i , v 
		}
	}
	fmt.Println(maxi,maxValue)

	//求和
	sum := 0
	for _,v := range arr3{
		sum += v 
	}
	fmt.Println(sum)

	printArray(&arr1)
	printArray(&arr3)
}


