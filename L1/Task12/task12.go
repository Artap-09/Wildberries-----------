package main

import "fmt"

func main() {
	var s = [5]string{"cat","cat","dog","cat","tree"}
	set:=make(map[string]struct{})

	for _,val:=range s {
		if _,ok:= set[val]; ok {
			continue
		}

		set[val]=struct{}{}
	}

	fmt.Println(set)
	
}