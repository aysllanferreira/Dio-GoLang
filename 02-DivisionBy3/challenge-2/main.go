package main

import "fmt"

func main() {
	for number := 1; number <= 100; number++ {
		switch {
		case number%3 == 0:
			fmt.Println("Pin")
		case number%5 == 0:
			fmt.Println("Pan")
		default:
			fmt.Println(number)
		}
	}
}
