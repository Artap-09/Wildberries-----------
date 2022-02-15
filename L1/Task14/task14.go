package main

import "fmt"

func main() {
	set:=make([]interface{},0)
	ch:=make(chan struct{})
	set=append(set, "sa",12,12.3,ch,true)

	for _,val:= range set{
		switch val.(type){
		case int:
			fmt.Print("int ")	
		case chan struct{}:
			fmt.Print("chan struct{} ")
		case bool:
			fmt.Print("bool ")
		case string:
			fmt.Print("string ")
		case float64:
			fmt.Print("float64 ")
		default:
			fmt.Println("Не могу определить тип")
		}

		fmt.Printf("%T\n",val)
	}
}