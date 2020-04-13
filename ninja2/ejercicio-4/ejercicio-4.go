package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("El valor de %v en decimal es %d , en binario es %b y en hexadecimal es %x\n" , a , a , a , a)

	b:= a << 1
	fmt.Printf("El valor de %v en decimal es %d , en binario es %b y en hexadecimal es %x\n" , b , b , b , b)
}
