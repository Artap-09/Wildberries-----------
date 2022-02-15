package main

import "fmt"

func main() {
	var x,y = 12,21
	x,y=y,x
	fmt.Println(x,y)
}