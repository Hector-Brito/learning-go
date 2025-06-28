package main

import "fmt"

func addOne(x * int){
	*x++
}

func freeMemory(x * int){
	*x = 0//Esto es innecesario.
	x = nil//Le asigna a la direccion del puntero nil
}

func main(){
	var number int = 10
	fmt.Println("El valor de la variable (numero) es:",number)
	addOne(&number)//Se pasa la direccion en memoria de la variable (number)
	fmt.Println("El valor de la variable (numero) ahora es:",number)
	freeMemory(&number)
	fmt.Println("El valor del la direccion en memoria de la variable (numero) ahora es nil, por lo que se libera espacio en memoria.")
	fmt.Println("El valor de la variable (numero) ahora es",number)
}