package main

import "fmt"

func main(){
	var num int

	if fmt.Scanf("%d",&num); num<0{
		fmt.Println("The Number is -ve")
	} else if num%2==0{
		fmt.Println(num,"is an even number")
	}else if num%2==1{
		fmt.Println(num,"is and odd number")
	}else {
		fmt.Println("That's a strange number");
	}
}
