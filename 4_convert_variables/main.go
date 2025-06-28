package main

import (
	"fmt"
)

func main() {
	var i int = 10

	var j float32 = 20
	//Convertir variables de un tipo a otro dependiendo de el tipo de dato de la variable a la que se asigna dicho valor (Esto solo funciona con tipos numericos).
	//i = int(j)//Aqui convertimos de float32 a int por que i es de tipo int.
	j = float32(i) //Aqui convertimos de int a float32 por que j es de tipo float32.
	fmt.Printf("El valor de 'i' es:%v y su tipo de dato es: %T",i,i)
	fmt.Println()
	fmt.Printf("El valor de 'j' es:%v y su tipo de dato es: %T",j,j)
}