package main

import "fmt"

func main() {
	min := 1
	max := 100

	for {
		guess := (min + max) / 2
		fmt.Println("Is your number", guess, "?")
		fmt.Print("Enter h (too high), l (too low), or c (correct): ")

		var s string
		fmt.Scan(&s)

		if s == "h" {
			// guess is too high
			max = guess - 1
		} else if s == "l" {
			// guess is too low
			min = guess + 1
		} else if s == "c" {
			// guessed correctly
			fmt.Println("Yay! I guessed your number!")
			break
		} else {
			fmt.Println("Please enter h, l, or c")
		}
	}
}
