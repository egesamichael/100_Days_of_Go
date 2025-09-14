package main

import "fmt"

func arrays() {

	// Arrays in Go
	var cars [3]string = [3]string{"BMW", "Toyota", "Honda"}
	fmt.Println("Cars:", cars)

	// Arrays have a fixed size we can't change it 
	// so we cant add more elements to it
	// cars[3] = "Ford" // This will give an error

	// We can also get the length of the array
	fmt.Println("Number of cars:", len(cars))

	// We can also use the shorthand syntax to declare and initialize an array
	fruits := [3]string{"Apple", "Banana", "Orange"}
	fmt.Println("Fruits:", fruits)

	// We can access elements of the array using index
	fmt.Println("First car:", cars[0])
	fmt.Println("Second fruit:", fruits[1])

	// We can also change the value of an element in the array
	cars[0] = "Mercedes"
	fmt.Println("Updated Cars:", cars)

}