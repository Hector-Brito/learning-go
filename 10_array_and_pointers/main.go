package main

import (
	"fmt"
)

func printArray(arr * [5]int){
	fmt.Println(*arr)
}

func main(){
	var my_arr [5]int = [5]int{1,2,3,4,5}
	var arr_pointer * [5]int = &my_arr
	printArray(arr_pointer)
}