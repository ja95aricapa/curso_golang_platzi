package main

import (
	"fmt"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgumentFunction(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValuesFunction(a int, b int) int {
	return a + b
}

func doubleReturn(a int) (c int, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgumentFunction(1, 2, "Hello World!")
	value := returnValuesFunction(3, 2)
	fmt.Println("Value:", value)
	value1, value2 := doubleReturn(2)
	fmt.Println("Value1:", value1, "Value2:", value2)
}
