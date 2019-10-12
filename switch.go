package main

import (
	"fmt"
)

func main(){
	var a int
	fmt.Scanln(&a)
	
	switch a{
	case 1:
		fmt.Println("Value is 1")
		break
	 case 2:
		fmt.Println("Value is 2")
		break
	default:
		fmt.Println("default case")
	}
	fmt.Println("Ended")
}