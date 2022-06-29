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
	X  = make(map[string]string)
)

type Line struct {
	CurrentLine int
	Value       string
}

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

func errFunc() {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 10000; j++ {
				mu.Lock()
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10
				mu.Unlock()
			}
		}()
	}
	for key, value := range m {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
	log.Print("done")
}

func main() {

	// Bài 1"
	fmt.Println("Bài 1:")
	chanRoutine()
	fmt.Println("-----------------------------------------------")

	// Bài 2:
	fmt.Println("Bài 2:")

	wg.Add(3)
	go func() {
		for i := 1; i <= 1000; i++ {
			mu.Lock()
			X["Key i "+strconv.Itoa(i)] = "Value " + strconv.Itoa(i)
			mu.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for j := 1; j <= 1000; j++ {
			mu.Lock()
			X["Key j "+strconv.Itoa(j)] = "Value " + strconv.Itoa(j)
			mu.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for k := 1; k <= 1000; k++ {
			mu.Lock()
			X["Key k "+strconv.Itoa(k)] = "Value " + strconv.Itoa(k)
			mu.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()

	// for key, value := range X {
	// 	fmt.Printf("%v: %v\n", key, value)
	// }

	fmt.Println("Length X: ", len(X))

	fmt.Println("-----------------------------------------------")

	// Bài 3:
	fmt.Println("Bài 3:")

	errFunc()

	fmt.Println("-----------------------------------------------")

	// Bài 4:

	fmt.Println("Bài 4:")

	data := make(chan string, 10)

	go worker(data)

	for v := range data {
		fmt.Println(v)
	}

	fmt.Println("Xong!")
}

func worker(data chan string) {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Print(err)
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		data <- line
	}
	close(data)

	file.Close()
}
