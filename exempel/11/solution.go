package main

import "fmt"

func sumAll(numbers []int) int {
	var sum int
	for _, value := range numbers {
		sum = sum + value
	}
	return sum
}

func countEven(numbers []int) int {
	var count int
	for _, value := range numbers {
		if value%2 == 0 {
			count = count + 1
		}
	}
	return count
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	sum := sumAll(numbers)
	even := countEven(numbers)
	fmt.Printf("sum=%d even=%d\n", sum, even)
}
