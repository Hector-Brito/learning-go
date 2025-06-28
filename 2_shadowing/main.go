package main

import (
	"fmt"
)

//Variable declarada fuera del main, es posible volverla a declarar dentro del main, cambiando tanto su valor como el tipo de dato
var i int = 1

func main(){
	i = 2//Se puede reasignar el valor✅
	i := "Soy string"//Se puede redeclarar✅
	
	//Variable declarada dentro del main, no puede ser redeclarada pero si puede cambiar su valor solo si pertenece al mismo tipo de dato.
	//Esto es porque no se pueden redeclarar variables dentro del mismo scope.
	var j int = 1
	//j:=2 Esto no ❌
	j= 3 // Esto si ✅
	fmt.Printf("El valor de 'i' es: %v y el tipo es:%T\n",i,i)
	fmt.Printf("El valor de 'j' es: %v y el tipo es:%T\n",j,j)
}