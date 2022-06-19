package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func chanRoutine() {
	wg.Add(1)
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		wg.Done()
	}()
	log.Print("hello 2")
	wg.Wait()
}

func main() {
	// Bài 1"
	fmt.Println("Bài 1:")
	chanRoutine()
	fmt.Println("-----------------------------------------------")

	// Bài 2:
	fmt.Println("Bài 2:")
	X := make(map[string]string)

	for i := 1; i <= 1000; i++ {
		X["Key "+strconv.Itoa(i)] = "Value " + strconv.Itoa(i)
	}
	for key, value := range X {
		log.Printf("%v: %v", key, value)
	}

	log.Print("Length: ", len(X))

}
