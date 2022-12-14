package main

import (
	pkg "curso_golang_platzi/src/mypackage"
	"fmt"
)

// NOTA: go mod init así solito, en la raíz de la carpeta del proyecto, para arreglar los problemas de importacion de paquetes
// el modulo se llama mypackage, pero se puede usar como alias pkg para renombrarlo

func main() {
	var myCar pkg.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)
	// var myCar2 pkg.carPrivate
	// myCar2.brand = "BMW"
	// myCar2.year = 2021
	// fmt.Println(myCar2)

	pkg.PrintMessage("Hola Platzi")
	// pkg.printMessage("Hola Platzi")

}
