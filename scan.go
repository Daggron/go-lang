package main

import "fmt"

func swapvar(a,b int)(int,int){
	return b,a
}

func main(){
	var a,b int
	fmt.Scanf("%d \n %d",&a,&b);
	a,b = swapvar(a,b)

	fmt.Println(a,b)
}
