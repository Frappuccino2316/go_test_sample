package main

import (
	"fizz/calculate"
	"fmt"
)

func main() {
	for i := 1; i < 31; i++ {
		if calculate.FizzBuzz(i) {
			fmt.Println("FizzBuzz!")
		} else if calculate.Fizz(i) {
			fmt.Println("Fizz")
		} else if calculate.Buzz(i) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
