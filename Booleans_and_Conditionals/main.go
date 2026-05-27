package main

import (
	"fmt"
)

func main() {
	// Booleans and Conditionals
	// age := 45

	// if age < 18 {
	// 	fmt.Println("You are a minor.")
	// } else if age >= 18 && age < 65 {
	// 	fmt.Println("You are an adult.")
	// } else {
	// 	fmt.Println("You are a senior.")
	// }

	// Using boolean and condition in picking scores and grading
	// fmt.Println("Enter your score:")
	// var score int
	// fmt.Scan(&score)

	// if score >= 90 {
	// 	fmt.Println("Grade: A")
	// } else if score >= 80 {
	// 	fmt.Println("Grade: B")
	// } else if score >= 70 {
	// 	fmt.Println("Grade: C")
	// } else if score >= 60 {
	// 	fmt.Println("Grade: D")
	// } else {
	// 	fmt.Println("Grade: E")
	// }

	//nesting an if statement inside a for loop in a slice of names with coninue and break statement
	names := []string{"Alice", "Bob", "Charlie", "Bato"}
	for index, name := range names {
		if name == "Charlie" {
			fmt.Printf("Found %s at index %d\n", name, index)
			continue // skip the rest of the loop and move to the next iteration
		}
		if index > 2 {
			fmt.Printf("Index %d is greater than 2, breaking the loop\n", index)
			break // exit the loop entirely
		}
	}
}
