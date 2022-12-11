package main

import (
	"fmt"
	"math"
)

func main3() {
	//PRIMITIVE DATA TYPES

	//ints : Integer
	{
		//int
		var a int = 1 //Depending on the OS it will assign 32 or 64 bits
		//int8
		var b int8 = -128 //8 bits (-128 to 127)
		//int16
		var c int16 = -int16(math.Pow(2, 15)) //16 bits (-2^15 to 2^15-1)
		//int32
		var d int32 = -int32(math.Pow(2, 31)) //32 bits (-2^31 to 2^31-1)
		//int64
		var e int64 = -int64(math.Pow(2, 63)) //64 bits (-2^63 to 2^63-1)
		fmt.Println(a, " ", b, " ", c, " ", d, " ", e)
	}

	//uint : Positive Integer
	{
		//uint
		var a uint //Depending on the OS it will assign 32 or 64 bits
		//uint8
		var b uint8 = uint8(math.Pow(2, 8)) - 1 //8 bits (0 to 2^8 - 1)
		//uint16
		var c uint16 = uint16(math.Pow(2, 16)) - 1 //16 bits (0 to 2^16 - 1)
		//int32
		var d uint32 = uint32(math.Pow(2, 32)) - 1 //32 bits (0 to 2^32 - 1)
		//uint64
		var e uint64 = uint64(math.Pow(2, 64)) - 1 //64 bits (0 to 2^64 - 1)
		fmt.Println(a, " ", b, " ", c, " ", d, " ", e)
	}

	//float: decimal numbers
	{
		var a float32 = 3.4e38  //32 bits - +/- 1.18e-38 to +/- 3.4e38
		var b float64 = 1.8e307 //64 bits - +/- 2.23e-308 to +/- 1.8e308
		fmt.Println(a, " ", b, " ")
	}

	//string: text
	{
		var a string = "Hello World" // string must be declared with "" and not with any other type of commas
		fmt.Println(a)
	}

	//bool: true or false
	{
		var a bool = true
		var b bool = false
		fmt.Println(a, " ", b, " ")
	}

	//complex: complex numbers
	{
		var a complex64 = 10 + 10i      //Real part and imaginary, both float32
		var b complex128 = 1000 + 2000i //Real part and imagniary, both float64
		fmt.Println(a, " ", b, " ")
	}
}
