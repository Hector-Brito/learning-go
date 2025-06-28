package main

import "fmt"

func main() {
	/*
		* Se utiliza en cualquier tipo de dato para decir que es tipo puntero.
		& Se utiliza para acceder a la direccion de memoria de una variable.
	*/
	//1 ejemplo
	var variable int = 15//La variable es una referencia del valor y direccion en memoria que manejamos utilizando el alto nivel del lenguaje.
	var pointer * int = &variable//Puntero que apunta a la direccion en memoria de la variable.
	fmt.Println("La variable tiene el valor:",variable)//output:15
	fmt.Println("El puntero tiene el valor:",pointer)//output:0xc00000a0d8
	fmt.Println("Cuando desreferenciamos el puntero tiene este valor:",*pointer)//output:15
	*pointer = 20 //Cambio el valor que hay en la direccion en memoria de la variable a 20.
	fmt.Println("Despues del cambio de valor desreferenciando el puntero, la variable tiene el siguiente valor:",variable)//output:20
	fmt.Println()
	//2 ejemplo
	var newPointer * int = new(int) //new genera un valor tipo puntero
	fmt.Println("El valor del puntero (newPointer) es:",newPointer)
	fmt.Println("El valor al que apunta el puntero (newPointer) es:",*newPointer)
	*newPointer = 5
	fmt.Println("El valor al que apunta el puntero (newPointer) ahora es:",*newPointer)

	newPointer = nil
	
}