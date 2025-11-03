package main


import (
	"fmt"
)

func main(){
	/*
	A pesar de ser un puntero de tipo int,
	al asignarle el valor nil(ausencia de valor)
	no apuntara a un lugar en memoria como tal.
	*/
	var int_null_pointer *int = nil
	fmt.Println(int_null_pointer)
}