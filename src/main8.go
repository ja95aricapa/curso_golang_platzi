package main

import (
	"fmt"
)

func main() {
	//Defer
	defer fmt.Println("Hola") // It is executed at the end of the code, it is recommended to use a maximum of one defers per function
	//In the case there is more than one defer, they will be treated as a stack (LIFO)
	defer fmt.Println("Que tal")
	fmt.Println("Mundo")

	for i := 0; i < 10; i++ {

		//continue
		if i == 2 {
			fmt.Println("Es 2")
			continue //It jumps to the next iteration without executing the code below
		}

		fmt.Println(i)

		if i == 8 {
			fmt.Println("Break")
			break //It breaks the loop without executing the next iterations
		}
	}

}
