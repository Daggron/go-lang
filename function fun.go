package main

import (
	"fmt"
)

func getAnswer(x int , y int) int  {
	return x+y
}

func main()  {
	var add  int  = getAnswer(10,20)

	fmt.Println("the sum is ",add);
}
