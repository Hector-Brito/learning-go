package main

import (
	"fmt"
)

func makeMeOlder(age int){
	age += 5
}

func main(){
	/*
	Al pasar la variable, crea una copia dentro de la funcion y maneja esa copia.
	*/
	var myAge int = 10
	makeMeOlder(myAge)
	fmt.Println(myAge)//myAge is still 10
}