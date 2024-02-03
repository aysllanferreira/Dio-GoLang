package main

import "fmt"

func main() {
	var number int = 1
	for number <= 100 {
		if number%3 == 0 {
			fmt.Println(number)
		}
		number++
	}
}