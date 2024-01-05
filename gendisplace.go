package main

import (
	"fmt"
)

func main() {
	a := make([]float64, 4)
	fmt.Println("Enter a value for acceleration: ")
	fmt.Scan(&a[0])
	fmt.Println("Enter a value for initial velocity: ")
	fmt.Scan(&a[1])
	fmt.Println("Enter a value for initial displacement: ")
	fmt.Scan(&a[2])
	fmt.Println("Enter a value for time: ")
	fmt.Scan(&a[3])
	fn := GenDisplaceFn(a[0], a[1], a[2])
	fmt.Println("Result:", fn(a[3]))
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 1/2*a*t*t + vo*t + so
	}
}
