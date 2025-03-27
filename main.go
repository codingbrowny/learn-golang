package main

import "fmt"

func main() {

	var scores [3]int = [3]int{20, 40, 60}
	fmt.Println(scores)

	//Access element in array
	fmt.Println(scores[0])

	//Change value at array index
	scores[0] = 100

	fmt.Println(scores)

	// Compiler Count ELements
	names := [...]string{"Sam", "Bob", "Phil", "Browny"}

	fmt.Println(names, len(names))

}