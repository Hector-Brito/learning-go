package main

import "fmt"

//Solo ser puede declarar de esta forma fuera del main, no se puede usando :=
var global_scope string = "Soy de global scope"

//Declarar multiples variables fuera  del main
var (
	var1 int = 20
	var2 int = 30
	var3 int = 40
)

func main(){
	fmt.Printf("El valor de 'global_scope' es:(%v), El tipo de 'global_scope' es:(%T)",global_scope,global_scope)
	//Declaracion de tipo de dato
	var flotante32 float32
	//Declaracion inferida con :=
	inference_var := 20
	//Declaracion directa
	var entero int = 10
	var flotante64 float64 = 20.
	fmt.Printf("El valor de 'entero' es:(%v), El tipo de 'entero' es:(%T)",entero,entero)
	fmt.Printf("El valor de 'flotante32' es:(%v), El tipo de 'flotante32' es:(%T)",flotante32,flotante32)
	fmt.Printf("El valor de 'flotante64' es:(%v), El tipo de 'flotante64' es:(%T)",flotante64,flotante64)
	fmt.Printf("El valor de 'inference_var' es:(%v), El tipo de 'inference_var' es:(%T)",inference_var,inference_var)
}