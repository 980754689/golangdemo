package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := arr[2:5]
	fmt.Println(s1)

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	updateSlice(s1)
	fmt.Println("updateSlice(s1) = ", s1)

	updateSlice(arr[:6])
	fmt.Println("updateSlice(arr[:6]) = ", arr[:6])

	updateSlice(arr[2:6])
	fmt.Println("updateSlice(arr[2:6]) = ", arr[2:6])

	updateSlice(arr[:])
	fmt.Println("updateSlice(arr[:]) = ", arr[:])

    arr2 := [...]int{1,2,3,4,5,6,7,8,9,10}
	s2 := arr2[2:6]
	s3 := s2[3:5]
	s4 := s2[3:7]
	s5 := s2[1:5]

	fmt.Println(s2,s3,s4,s5)

}
