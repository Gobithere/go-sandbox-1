package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	ID   int64
	Name string
}

var productChan = make(chan Product)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	go getProductDetailsFromExternal(1)
	select {
	case product := <-productChan:
		fmt.Println(product)
	case <-ctx.Done():
		fmt.Println("Timeout")
	}
	fmt.Println("End of main function")
}
func getProductDetailsFromExternal(productId int64) {
	time.Sleep(10 * time.Second)

	productChan <- Product{
		ID:   productId,
		Name: "Çamaşır Makinesi",
	}
}
