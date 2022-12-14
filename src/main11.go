package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main11() {

	fmt.Println()
	fmt.Println("=======")
	fmt.Println("STRUCTS")
	fmt.Println("=======")
	fmt.Println()

	fmt.Println("En otros lenguajes existe la estructura llamada Clase, donde se definen atributos y métodos. ")
	fmt.Println("En Go no existe esta estructura pero ofrece 'structs' que según tiene mas poder.")
	fmt.Println()

	fmt.Println("Declaración:")
	fmt.Println("------------")
	fmt.Println()

	fmt.Println("Se declara como otro bloque de función, pero cambiando func por 'type' y luego del nombre se coloca 'struct'")
	fmt.Println("type car struct { ... }")
	fmt.Println()

	fmt.Println("se declaran sus atributos como pares nombre-tipo, uno por línea, tal cual como en el código de este archivo.")
	fmt.Println("atributico string")
	fmt.Println()

	fmt.Println("Primera forma de instanciación:::")
	fmt.Println("---------------------------------")
	fmt.Println()

	myCar := car{brand: "Renault", year: 2021}

	fmt.Println("myCar := car{brand: \"Renault\", year: 2021}")
	fmt.Println()

	fmt.Println("FMT Imprime: ")
	fmt.Println(myCar)
	fmt.Println()

	fmt.Println("Otra forma de imprimirlo en consola:")
	fmt.Printf("Structsito = %+v", myCar)
	fmt.Println()

	fmt.Println()
	fmt.Println("Instanciación como una clase vacía:")
	fmt.Println("-----------------------------------")
	fmt.Println("")

	fmt.Println("Se intancia así: ")
	fmt.Println("var otherCar car")
	fmt.Println()

	fmt.Println("Y se va asignando los valores de cada atributo:")
	fmt.Println("otherCar.brand = \"Chevrolet\"")
	fmt.Println()

	var otherCar car
	otherCar.brand = "Chevrolet"

	fmt.Println("Si un atributo del struct no se inicializa, toma un Zero Value, y por lo tanto se imprime así lo referente al año:")

	fmt.Println(otherCar)
	fmt.Println()

	fmt.Println()
}
