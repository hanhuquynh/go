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

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetAge() int64 {
	return u.age
}

func (u *User) GetGender() bool {
	return u.gender
}

func (u *User) GetAddress() string {
	return u.address
}

// Bài 6
func bai6() {
	num := []int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	sort.Ints(num)
	num2 := num[1:7]
	fmt.Println("Bài 6:")
	fmt.Printf("%v\n", num)
	fmt.Printf("%v\n", num2)
	fmt.Println("----------------------------------------------------------------------")
	/*
		num3 := num[1:15]
		fmt.Printf("%v\n", num3)
		=> Có lỗi do capacity của slice = 10
	*/
}

func main() {
	// Bài 6
	bai6()

	// Bài 7
	fmt.Println("Bài 7: ")
	user := User{name: "Hà Như Quỳnh", age: 21, gender: true, address: "Bắc Ninh"}
	fmt.Println("Name:", user.GetName())
	fmt.Println("Age:", user.GetAge())
	if user.GetGender() == true {
		fmt.Println("Gender: Male")
	} else {
		fmt.Println("Gender: Female")

	}
	fmt.Println("Address:", user.GetAddress())
}
