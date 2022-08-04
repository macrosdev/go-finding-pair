package main

import (
	"fmt"
	"sort"
)

func main() {
	length := 0
	fmt.Println("Enter the number of inputs")
	fmt.Scanln(&length)
	fmt.Println("Enter the inputs")
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanln(&numbers[i])
	}
	fmt.Println("Enter the Target")
	target := 0
	fmt.Scanln(&target)

	sort.Ints(numbers)
	for i := 0; i < length; i++ {
		rest := target - numbers[i]
		for j := i + 1; j < length; j++ {
			if rest < numbers[j] {
				break
			}
			if rest == numbers[j] {
				fmt.Println(numbers[i], rest)
				return
			}
		}
	}
	fmt.Println("No Pair")
}
