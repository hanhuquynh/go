package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
)

type Line struct {
	CurrentLine int
	Value       string
}

func worker(jobs <-chan string, results chan<- string) {
	for v := range jobs {
		results <- v
	}
}

func main() {
	size := 10
	jobs := make(chan string, size)
	results := make(chan string, size)

	file, err := os.Open("file.txt")

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		jobs <- fileScanner.Text()
	}
	close(jobs)

	file.Close()

	for i := 0; i < 3; i++ {
		go worker(jobs, results)
	}

	fmt.Println("Num goroutine:", runtime.NumGoroutine())

	for i := 1; i <= size; i++ {
		rs := Line{
			CurrentLine: i,
			Value:       <-results,
		}
		fmt.Printf("%v. %v\n", rs.CurrentLine, rs.Value)
	}

}
