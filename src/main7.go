package main

import "fmt"

var user, password = "abcd123", "123456"

// reto 1
func paroImpart(number int) string {
	if number%2 == 0 {
		return "par"
	} else {
		return "impar"
	}
}

// reto 2
func loggin(usuario, contraseña string) bool {
	if (usuario == user) && (contraseña == password) {
		return true
	} else {
		return false
	}
}

func main7() {
	reto1 := paroImpart(12)
	fmt.Println(reto1)
	reto2 := loggin("abcd123", "123456")
	fmt.Println(reto2)
}
