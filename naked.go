package main

import "fmt"

func getData()(x,y int){
	 x=6
	 y=7
	return
}

func main(){
	var a,b  = getData()
	fmt.Println(a,b)
}