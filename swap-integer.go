package main

import "fmt"

func swapInt(a,b int) (int, int){
	return b,a
}

func main() {
	var a,b int
	 fmt.Scanf("%d \n %d",&a,&b)
	a , b = swapInt(a,b)
	fmt.Println("swap" ,a,b)

}
