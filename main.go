package main

// func sum(a, b int) (result int) {
// 	result = a + b
// 	return result
// }

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

	// Slice

	// s1 := []int{2, 4, 5, 6, 7, 8, 10, 13}
	// s2 := s1[:]   // lấy hết các phần tử của a
	// s3 := s1[3:]  //index 3 đến hết
	// s4 := s1[:6]  // từ đầu đến index 5
	// s5 := s1[3:6] // từ index 3 đến index 5

	// fmt.Printf("s1 %v, %v, %v\n", s1, len(s1), cap(s1))
	// fmt.Printf("s2 %v, %v, %v\n", s2, len(s2), cap(s2))
	// fmt.Printf("s3 %v, %v, %v\n", s3, len(s3), cap(s3))
	// fmt.Printf("s4 %v, %v, %v\n", s4, len(s4), cap(s4))
	// fmt.Printf("s5 %v, %v, %v\n", s5, len(s5), cap(s5))

	// s6 := make([]int, 10, 20)
	// fmt.Printf("s6 %v, %v, %v\n", s6, len(s6), cap(s6))
	// s6 = append(s6, 20, 55, 62, 45, 20, 55, 62, 45, 20, 55)
	// fmt.Printf("s6 %v, %v, %v\n", s6, len(s6), cap(s6))

	// s7 := []int{1, 2, 3}
	// s8 := []int{4, 5, 6}

	// s9 := append(s7, s8...)
	// fmt.Printf("s9 %v, %v, %v\n", s9, len(s9), cap(s9))

	// Condition

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

	// familyName("Liam", 5)
	// familyName("Jenny", 14)
	// familyName("Anja", 30)

	// result := sum(3, 5)
	// fmt.Println("a + b =", result)

	// a, name := test(5, "Quỳnh")
	// fmt.Println(a, name)

	// var person1 Person

	// person1.name = "Quỳnh"
	// person1.address = "Bắc Ninh"
	// person1.age = 21

	// fmt.Println(person1.name)
	// fmt.Println(person1.age)
	// fmt.Println(person1.address)

	// printPerson(person1)

	// Map
	// student := map[string]string{
	// 	"A": "Bắc Ninh",
	// 	"B": "Hà Nội",
	// 	"C": "Bắc Giang",
	// }

	// student2 := student

	// student2["A"] = "BN"

	// fmt.Println(student)
	// fmt.Println(student2)

}

// Struct

// type Person struct {
// 	name    string
// 	age     int
// 	address string
// }

// func printPerson(per Person) {
// 	fmt.Println(per.name)
// 	fmt.Println(per.age)
// 	fmt.Println(per.address)
// }

// func test(a int, name string) (rs1 int, rs2 string) {
// 	rs1 = a * a
// 	rs2 = "My name is " + name
// 	return
// }
