package main

import (
	"fmt"
)

func main(){
	/*
	Al utilizar el puntero de y en x podemos cambiar el valor de y.
	*/
	var y int = 3
	var x *int = &y
	*x = 10
	fmt.Println("El valor de y es",y) 
}