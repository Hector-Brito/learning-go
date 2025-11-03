package main

import (
	"fmt"
)

func add(x,y int)int{
	return x+y
}

func main(){
	var function_pointer func(int,int) int = add
	fmt.Println(function_pointer(5,5))
	/*
	Mejores practicas para usar punteros en Go.
	1- Evitar usar unsafe.Pointer a menos que sea absolutamente necesario.
	2- Evitar usar aritmetica de punteros a menos que sea necesario.
	3- Usar "null pointers" para indicar uninitialized or deallocated memory.
	4- Usar "function pointers" con moderacion
	5- Usar punteros para reducir el uso de memoria, al utilizar punteros para almacenar grandes estructuras de datos, es posible evitar copiar los datos enteros cada vez que este es pasado a una funcion o almacenado en una estructura de datos.
	
	*/

}