package main


import (
	"fmt"
)


func swap(x *int, y *int){
	temp := *x
	*x = *y
	*y = temp
}

func double(x *int) *int{
	result := *x *2
	return  &result
}

func main(){
	var x int = 1
	var y int = 2
	fmt.Println(x,y)
	swap(&x,&y)
	fmt.Println(x,y)
	var pointer *int = double(&x)
	fmt.Println(*pointer)

}