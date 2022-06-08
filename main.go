package main

import "fmt"

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func sum(a, b int) (result int) {
	result = a + b
	return result
}

func main() {
	// Variables
	// var a, b int = 1, 5

	// fmt.Println(a + b)

	// var name = "Quynh"
	// fmt.Println("My name is", name)

	// var c = true
	// fmt.Println(c)

	// d := 10 // = var d int = 10
	// e := 20

	// fmt.Println(d + e)

	// age := 18

	// fmt.Println(age)

	// var f = "hello world"
	// fmt.Printf("Type is: %T \n", f)

	// var g = 100
	// fmt.Printf("%v%% \n", g)

	// Array

	// var arr = [4]int{1, 2, 3, 4}
	// arr2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(arr)
	// fmt.Println(arr2)

	// var arr3 = [...]string{"red", "green", "blue", "yellow", "pink"}
	// arr4 := [...]string{"A", "B", "C", "D"}
	// fmt.Println(arr3[3])
	// arr3[3] = "black"
	// fmt.Println(arr3)
	// fmt.Println(arr4)

	// arr5 := [5]int{}
	// fmt.Println(arr5)
	// arr3[2] = "gray"
	// fmt.Println(arr3)

	// arr6 := [5]int{1: 15, 3: 20} // gán 15 vào index thứ 1, gán 20 vào index thứ 3
	// fmt.Println(arr6)

	// fmt.Println("Length arr3:", len(arr3))
	// fmt.Println("Length arr4:", len(arr4))

	// a := 14
	// b := 14
	// if a < b {
	// 	fmt.Println("a is less than b.")
	// } else if a > b {
	// 	fmt.Println("a is more than b.")
	// } else {
	// 	fmt.Println("a and b are equal.")
	// }

	// day := 3

	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// case 7:
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Not a weekday")
	// }

	// for i := 0; i <= 100; i++ {
	// 	if i%2 != 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// for i := 0; i < 100; i += 10 {
	// 	if i == 20 {
	// 		continue
	// 	}
	// 	if i == 50 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	familyName("Liam", 5)
	familyName("Jenny", 14)
	familyName("Anja", 30)

	result := sum(3, 5)
	fmt.Println("a + b =", result)

	a, name := test(5, "Quỳnh")
	fmt.Println(a, name)

	var person1 Person

	person1.name = "Quỳnh"
	person1.address = "Bắc Ninh"
	person1.age = 21

	fmt.Println(person1.name)
	fmt.Println(person1.age)
	fmt.Println(person1.address)

	printPerson(person1)

}

type Person struct {
	name    string
	age     int
	address string
}

func printPerson(per Person) {
	fmt.Println(per.name)
	fmt.Println(per.age)
	fmt.Println(per.address)
}

func test(a int, name string) (rs1 int, rs2 string) {
	rs1 = a * a
	rs2 = "My name is " + name
	return
}
