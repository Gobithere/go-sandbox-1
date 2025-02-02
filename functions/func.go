package main

import "fmt"

func main() {

	// function as parameter value

	numbers := []int{1, 2, 3, 4, 5}

	doubledNumbers := transformNumbers(&numbers, double)
	tripleNumbers := transformNumbers(&numbers, triple)

	fmt.Println(doubledNumbers)
	fmt.Println(tripleNumbers)

}

type transformfn func(int) int

func transformNumbers(numbers *[]int, transform transformfn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {

		dNumbers = append(dNumbers, transform(value))

	}
	return dNumbers
}

// Feleeen.
//
// double functions
func double(number int) int {
	return number * 2
}
func triple(number int) int {
	return number * 3
}
