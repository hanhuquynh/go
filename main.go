package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

// Bài 2
type Serve struct {
	Name  string
	Class string
}

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

func main() {
	// Bài 2
	fmt.Println("Bài 2: ")
	content, err := ioutil.ReadFile("serve.json")
	if err != nil {
		log.Fatal(err)
	}

	var serve []Serve
	err = json.Unmarshal(content, &serve)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", serve)
	fmt.Println("----------------------------------------------------------------------")

	// Bài 3
	fmt.Println("Bài 3: ")
	for _, s := range serve {
		classObj := s.Class
		classObj = strings.ToLower(classObj)

		if strings.Contains(classObj, "admin") {
			log.Println(classObj)
		}
	}
	fmt.Println("----------------------------------------------------------------------")
	// Bài 4
	fmt.Println("Bài 4:")

	obj := Serve{
		Name:  "fileCustome",
		Class: "org.cofax.cds.FileServlet.Custome",
	}

	serve = append(serve, obj)
	fmt.Printf("%v\n", serve)

	fmt.Println("----------------------------------------------------------------------")

	// Bài 5
	fmt.Println("Bài 5:")

	for i := 0; i < len(serve); i++ {
		fmt.Printf("%p - %v\n", &serve[i], serve[i])
	}

	fmt.Println("----------------------------------------------------------------------")

	// Bài 6
	fmt.Println("Bài 6:")
	num := []int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	sort.Ints(num)
	fmt.Printf("%v\n", num)
	num2 := num[1:7]
	fmt.Printf("%v\n", num2)
	/*
		num3 := num[1:15]
		fmt.Printf("%v\n", num3)
		=> Có lỗi do capacity của slice = 10
	*/
	fmt.Println("----------------------------------------------------------------------")

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
