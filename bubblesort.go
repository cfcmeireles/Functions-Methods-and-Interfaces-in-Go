package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter up to 10 integers: ")
	var integers []int
	var userInput int
	for i := 0; i < 10; i++ {
		fmt.Scan(&userInput)
		integers = append(integers, userInput)
		if len(integers) <= 10 {
			BubbleSort(integers)
		} else {
			fmt.Println("Error, please insert up to 10 integers")
		}
	}
}

func Swap(sli []int, i int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-1; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		Swap(sli, i)
	}
	fmt.Println("Sorted:", sli)
}
