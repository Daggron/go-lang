package main

import "fmt"


func swap(a , b string) (string,string){
	return b,a
}
func main(){
	var  a,b = swap("Hello ","World")
	 fmt.Println(a,b)

}
