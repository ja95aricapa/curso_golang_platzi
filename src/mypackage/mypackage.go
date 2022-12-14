package mypackage

import "fmt"

// CarPublic Car con acceso publico (primera letra en mayuscula)
type CarPublic struct {
	Brand string
	Year  int
}

// CarPublic Car con acceso privado (primera letra en minuscula)
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage Imprime un mensaje en consola
func PrintMessage(text string) {
	fmt.Println(text)
}

// printMessage Imprime un mensaje en consola
func printMessage(text string) {
	fmt.Println(text)
}
