package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

//var wg sync.WaitGroup

func getObject() {
	objects, _ := http.Get("https://api.restful-api.dev/objects/7")

	rawData, _ := io.ReadAll(objects.Body)

	type object struct {
		ID   string                 `json:"id"`
		Name string                 `json:"name"`
		Data map[string]interface{} `json:"data"`
	}

	objectData := object{}

	json.Unmarshal(rawData, &objectData)

	fmt.Println(objectData.Name)

	// fmt.Println(objectDataMap["CPU model"])
}

func getObject2() {
	time.Sleep(5 * time.Second)
	objects, _ := http.Get("https://api.restful-api.dev/objects/8")

	fmt.Println(objects.StatusCode)
	rawData, _ := io.ReadAll(objects.Body)

	type object struct {
		ID   string                 `json:"id"`
		Name string                 `json:"name"`
		Data map[string]interface{} `json:"data"`
	}

	objectData := object{}

	json.Unmarshal(rawData, &objectData)

	fmt.Println(objectData.Name)

	// fmt.Println(objectDataMap["CPU model"])
}

func main() {
	getObject2()
	//Interval goroutine
	// interval := time.NewTicker(1 * time.Second)
	// for range interval.C {
	// 	go getObject()
	// 	go getObject2()
	// }

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
	// mychannel := make(chan string)

	// done := make(chan bool)
	// go func() {
	// 	message := <-mychannel
	// 	fmt.Println(message)
	// 	done <- true
	// }()
	// go func() {
	// 	mychannel <- "Hello"
	// }()

	//done channel pattern true dönene kadar main function bekler. wg kullanmak yerine done channel kullanılabilir.
	// <-done
	// fmt.Println("End of main function")
}
