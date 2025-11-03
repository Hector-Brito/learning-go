package main

import (
	"fmt"
)

func main(){
	var name string = "Simon"
	fmt.Printf("Address of '%v' is %v",name,&name)
}