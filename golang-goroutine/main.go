package main

import (
	"fmt"
	"sync"
)

// func printX() {
// 	for i := 0; i < 20; i++ {
// 		fmt.Print("X")
// 	}
// 	wg.Done()
// }
// func printY() {
// 	for i := 0; i < 20; i++ {
// 		fmt.Print("Y")
// 	}
// 	wg.Done()
// }

var wg sync.WaitGroup

func main() {
	// wg.Add(2)
	// go printX()
	// fmt.Println("")
	// go printY()
	// wg.Wait()
	// for {

	// 	listener := time.NewTicker(10 * time.Second)
	// 	select {
	// 	case <-listener.C:
	// 		fmt.Println("hello")
	// 	}

	// }

	//channels
	mychannel := make(chan string)

	done := make(chan bool)
	go func() {
		message := <-mychannel
		fmt.Println(message)
		done <- true
	}()
	go func() {
		mychannel <- "Hello"
	}()

	//done channel pattern true dönene kadar main function bekler. wg kullanmak yerine done channel kullanılabilir.
	<-done
	fmt.Println("End of main function")
}
