package main

import (
	"log"
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
	chanRoutine()
}
