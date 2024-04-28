package main

import "fmt"

func main(){
	// s := []int{2, 3, 5, 7, 11, 13}
	// printSlice(s)
	// // len = 6, cap = 6

	// // Slice the slice to give it zero length.
	// s = s[:0]
	// printSlice(s)
	// // len = 0, cap = 6

	// // Extend its length.
	// s = s[:4]
	// printSlice(s)
	// // len = 4, cap = 6

	// // Drop its first two values.
	// s = s[2:]
	// printSlice(s)
	// // len = 2, cap = 4
	// //
	//  var s1 []int
	// 	fmt.Println(s1, len(s1), cap(s1))
	// 	if s1 == nil {
	// 		fmt.Println("nil!")
	// 	}

	// b := make([]int, 5)
	// // printSlice(b)
	// c := b
	// d := c[:2]
	// d[1] = 100
	// printSlice(b)
	// printSlice(c)
	// printSlice(d)


	// append
	a := make([]int, 2)
	b := make([]int, 3)
	b[0] = 11
	a = append(a, 12, 13, 14)
	printSlice(a)
	printSlice(b)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
