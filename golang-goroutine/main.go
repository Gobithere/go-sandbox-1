package main

import (
	"fmt"
	"sync"
	"time"
)

func printX() {
	for i := 0; i < 20; i++ {
		fmt.Print("X")
	}
	wg.Done()
}
func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	// wg.Add(2)
	// go printX()
	// fmt.Println("")
	// go printY()
	// wg.Wait()
	for {

		listener := time.NewTicker(10 * time.Second)
		select {
		case <-listener.C:
			fmt.Println("hello")
		}

	}

}
