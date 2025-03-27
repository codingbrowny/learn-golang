package main

import "fmt"

func main() {

	age := 25
	name := "Browny"

	//PRINT
	fmt.Print("First Print")

	//NEW LINE
	fmt.Println("Printed on a new line")
	fmt.Println("Hello Go!")
	fmt.Println("My name is", name, "and I am", age, "years old")


	//FORMATTED STRINGS - (%_ FORMAT SPECIFIER)
	fmt.Printf("My name is %v and I am %v years old \n", name, age)
	fmt.Printf("%v is of type %T \n", name, name)

	//SAVE FORMATTED STRING
	var formattedStr = fmt.Sprintf("My name is %v and I am %v years old \n", name, age)
	fmt.Println(formattedStr)

}