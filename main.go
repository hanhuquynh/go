package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

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

	wg.Add(3)

	go func() {
		mu.Lock()
		for i := 1; i <= 1000; i++ {
			X["Key i "+strconv.Itoa(i)] = "Value " + strconv.Itoa(i)
		}
		mu.Unlock()
		wg.Done()
	}()

	go func() {
		mu.Lock()
		for j := 1; j <= 1000; j++ {
			X["Key j "+strconv.Itoa(j)] = "Value " + strconv.Itoa(j)
		}
		mu.Unlock()
		wg.Done()
	}()

	go func() {
		mu.Lock()
		for k := 1; k <= 1000; k++ {
			X["Key k "+strconv.Itoa(k)] = "Value " + strconv.Itoa(k)
		}
		mu.Unlock()
		wg.Done()
	}()
	wg.Wait()

	for key, value := range X {
		fmt.Printf("%v: %v\n", key, value)
	}

	fmt.Println("Length X: ", len(X))

	fmt.Println("-----------------------------------------------") 

	// Bài 4:

	fmt.Println("Bài 4:")

	file, err := os.Open("file.txt")
	if err != nil {
		log.Print(err)
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	file.Close()

}
