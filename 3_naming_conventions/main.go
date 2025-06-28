package main

import (
	"fmt"
)
/*
Naming conventions
✅La primera letra de una variable en Go determina su acceso y visibilidad.
✔️Mayuscula la primera letra es exportada y accesible desde otros paquetes
✔️Minuscula la primera letra es interna al paquete donde se define y no desde otros paquetes
✅Usar nombres de variables respecto a su vida util.
✔️Para contadores por ejemplo podemos usar: i, j, k, n, c.
ya que vamos a usarlas para iterar
✔️Para operaciones donde sabemos que la variable tendra mas vida util, usar nombres tan largos como su ciclo de vida. 
Por ejemplo podemos usar:numberSeason para tener el numero de una temporada,etc.
✅Cuando usemos acronimos colocarlos en mayusculas, mejora la lectura. 
✔️Por ejemplo:theURL, theHTTP, etc
*/
var Im string = "Yo me exporto"
var im string = "Yo solo soy accesible internamente."
func main(){
	fmt.Printf("Hola soy: %v",Im)
}