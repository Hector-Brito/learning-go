package main

import (
	"fmt"
	"unsafe"
)

func main(){
	/*
	Este puntero puede apuntar a cualquier tipo de dato
	*/
	var void_pointer unsafe.Pointer;
	fmt.Println(void_pointer)
}