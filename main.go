package main

import (
	"context"
	"fmt"
	"time"
)

func x(ctx context.Context) context.Context {
	return context.WithValue(context.Background(), "current", time.Now().UnixNano())
}

func main() {
	// Bài 1:
	fmt.Println("Bài 1:")
	time.Sleep(time.Second * 3)
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
	fmt.Println("Today:", today.Day())

	fmt.Println("--------------------------------------------")

	// Bài 3:
	fmt.Println("Bài 3:")

	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	go func(ctx context.Context) {
		for {
			time.Sleep(time.Second * 3)
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled:", ctx.Err())
				return
			default:
				fmt.Println("Sleep")
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)

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

	//	Bài 6:

	/*
		Thời gian dạng số được sử dụng với mốc đơn vị:
			-> second
			-> milliseconds
			-> microseconds
			-> nanoseconds
	*/

	// Bài 7:
	fmt.Println("Bài 7:")

	ctx2 := context.Background()
	ctx2 = x(ctx2)

	currentTime := ctx2.Value("current").(int64)
	fmt.Println(currentTime)

	time.Sleep(time.Second * 3)
	currentTimeAfter3s := time.Now().UnixNano()
	fmt.Println(currentTimeAfter3s)

	result := currentTimeAfter3s - currentTime
	fmt.Println(result)

	fmt.Println("--------------------------------------------")

	//	Bài 9:
	fmt.Println("Bài 9: ")

	time.AfterFunc(time.Millisecond*100, func() {
		fmt.Println("i'm study")
	})
	time.Sleep(time.Second)

	fmt.Println("--------------------------------------------")

	//	Bài 8:
	fmt.Println("Bài 8:")
	for {
		fmt.Println(time.Now().Unix(), "done")
		time.Sleep(time.Millisecond * 100)
	}
}
