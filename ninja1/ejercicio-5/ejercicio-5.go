package main

import "fmt"

type marc int

var x marc
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("El tipo de x es %T \n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("El tipo de y es %T \n", y)
}
