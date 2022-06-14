package main

import (
	"fmt"
	"time"
)

func main() {
	// Bài 1:
	fmt.Println("Bài 1:")
	for i := 1; i <= 3; i++ {
		timeNow := time.Now().UnixMilli()
		fmt.Println(timeNow)
		time.Sleep(time.Second * 3)
	}
	fmt.Println("Kết thúc")
	fmt.Println("--------------------------------------------")

	// Bài 2:
	fmt.Println("Bài 2:")

	today := time.Now()
	fmt.Println("Today:", today.Day(), today)

	fmt.Println("--------------------------------------------")

	// Bài 4:
	fmt.Println("Bài 4: ")
	m := time.Unix(0, 1592190294764144364)
	fmt.Println("Number_of_minutes:", m.Minute())
	fmt.Println("--------------------------------------------")

	//	Bài 5:
	fmt.Println("Bài 5:")
	d := time.Unix(1592190385, 0)
	fmt.Println("Day of the week:", d.Weekday())
	fmt.Println("--------------------------------------------")

	//	 Bài 6:

	/*
		Thời gian dạng số được sử dụng với mốc đơn vị:
			-> second
			-> milliseconds
			-> microseconds
			-> nanoseconds
	*/

	//	Bài 9:
	fmt.Println("Bài 9: ")

	f := func() {
		fmt.Println("I'm study")
	}
	time.AfterFunc(time.Millisecond*100, f)
	time.Sleep(time.Second)

	fmt.Println("--------------------------------------------")

	//	Bài 8:
	fmt.Println("Bài 8:")
	for {
		fmt.Println(time.Now().Unix(), "done")
		time.Sleep(time.Millisecond * 100)
	}
}
