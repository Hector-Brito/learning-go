package main

import (
	"fmt"
	//"unsafe"
)

func main(){
	var arr [3]int = [3]int{4,5,6}//Declarar una lista de 3 elementos de tipo numero.
	fmt.Println(arr)//Imprimo la lista

	var pointer * int = &arr[0]
	fmt.Println("Primer elemento del arr:",*pointer)
	*pointer++
	fmt.Println("El primer elemento del arr aumentado en 1:",*pointer)

	/*
		Incrementar el punto con aritmetica de punteros.
	*/
	//Es posible incrementar un puntero utilizando el paquete unsafe
	//pointer = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(pointer)) + uintptr(unsafe.Sizeof(arr[0]))))
	//fmt.Println(*pointer)
}