package main

import "fmt"

func main6() {

	//For conditional
	fmt.Printf("For conditional\n")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// For while
	fmt.Printf("For while\n")
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Printf("\n")

	// For forever
	fmt.Printf("For forever\n")
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

}
