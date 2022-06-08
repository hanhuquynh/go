package main

import (
	"fmt"
	"sort"
)

// Bài 7

type User struct {
	name    string
	age     int64
	gender  bool
	address string
}

func bai6() {
	num := []int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	sort.Ints(num)
	num2 := num[1:7]
	fmt.Printf("%v\n", num)
	fmt.Printf("%v\n", num2)
	/*
		num3 := num[1:15]
		fmt.Printf("%v\n", num3)
		=> Có lỗi do capacity của slice = 10
	*/
}

func main() {
	bai6()
}
