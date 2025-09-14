package main

import "fmt"

func variables() {
	// Variables in Go
	// Go is a statically typed language
	// We can't declare a variable with a type and then assign a value of a different type to it
	var name string = "John"
	var age int = 30
	var isStudent bool = false

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)

	// We can also use the shorthand syntax to declare and initialize a variable
	city := "New York"
	fmt.Println("City:", city)

	// Constants in Go
	const pi = 3.14
	fmt.Println("Value of Pi:", pi)

	// Data Types in Go
	var price float64 = 19.99
	var isAvailable bool = true
	var initial rune = 'A' // rune is an alias for int32, used to represent a Unicode code point

	fmt.Println("Price:", price)
	fmt.Println("Is Available:", isAvailable)
	fmt.Println("Initial:", initial)

	// Type Conversion in Go
	var x int = 10
	var y float64 = float64(x) // Convert int to float64
	fmt.Println("Converted Value:", y)

	var z int = int(y) // Convert float64 to int
	fmt.Println("Converted Value:", z)

	// Multiple Variable Declaration
	var a, b, c int = 1, 2, 3
	fmt.Println("Values:", a, b, c)

	// Underscore (_) can be used to ignore a value
	_, d := 5, 10
	fmt.Println("Value of d:", d)
}	