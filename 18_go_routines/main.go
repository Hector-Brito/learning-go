package main

import (
	"fmt"
	"time"
)

func main(){
	go func(name string) {
		fmt.Println("Hello",name)
	}("Juan")
	go func (name string)  {
		fmt.Println("Bye",name)
	}("Juan")
	fmt.Println("Main function")
	time.Sleep(time.Second)
}