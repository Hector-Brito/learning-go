package main

import (
	"fmt"
	"strconv"
)

func main(){
	var i int = 100
	k := strconv.Itoa(i)//Convertir de int a string
	j,err := strconv.Atoi(k)
	fmt.Printf("El valor de 'i' es:%v y su tipo de dato es: %T\n",i,i)
	fmt.Printf("El valor de 'k' es:%v y su tipo de dato es: %T\n",k,k)
	fmt.Printf("El valor de 'k' es:%v y su tipo de dato es: %T\n",j,j)
	fmt.Printf("El error es:%v\n",err)

}