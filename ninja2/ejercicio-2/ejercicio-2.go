package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (10 <= 22)
	c := (15 >= 12)
	d := (11 != 21)
	e := (4 < 11)
	f := (13 > 7)

	fmt.Printf("%v\t\t%v\t\t%v\n%v\t\t%v\t\t%v\n" , a , b , c , d , e , f)
}
