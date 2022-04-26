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

	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := arr2[2:6]
	s3 := s2[3:5]
	s4 := s2[3:7]
	s5 := s2[1:5]

	fmt.Println(s2, s3, s4, s5)

	arr2[0], arr2[1] = 100, 200
	fmt.Println(arr2)
	s6 := arr2[1:6]
	s7 := s6[3:7]
	s8 := s6[:7]
	fmt.Println(s6, s7, s8)
	fmt.Printf("s6 = %v len(s6) = %d cap(s6) = %d\n", s6, len(s6), cap(s6))
	fmt.Printf("s7 = %v len(s7) = %d cap(s7) = %d\n", s7, len(s7), cap(s7))
	fmt.Printf("s8 = %v len(s8) = %d cap(s8) = %d\n", s8, len(s8), cap(s8))

	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s12 := arr3[:5]
	s9 := append(s12, 11)
	s10 := append(s9, 12)
	s11 := append(s10, 13)

	fmt.Println(s9, s10, s11)
	fmt.Printf("s12 = %v arr3 = %v\n", s12, arr3)

}
