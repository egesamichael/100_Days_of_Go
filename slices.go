package main

import "fmt"

func slices(){

	// Slices in Go
	fmt.Println("Slices in Go")

	var scores = []int{90, 80, 70, 60, 50}
	fmt.Println("Scores:", scores)

	// Slices are dynamic in size we can add or remove elements from it
	scores = append(scores, 40)
	fmt.Println("Updated Scores:", scores)

	// We can also get the length of the slice
	fmt.Println("Number of scores:", len(scores))

	// We can access elements of the slice using index
	fmt.Println("First score:", scores[0])
	fmt.Println("Second score:", scores[1])

	// We can also change the value of an element in the slice
	scores[0] = 100
	fmt.Println("Updated Scores:", scores)

	// We can create a slice from an array
	var fruits = [5]string{"Apple", "Banana", "Orange", "Mango", "Grapes"}
	var fruitSlice = fruits[1:4] // This will create a slice with elements from index 1 to 3
	fmt.Println("Fruit Slice:", fruitSlice)

}