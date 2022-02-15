package main

import (
	"fmt"
	"strings"
)

func main(){
	var s string
	fmt.Scan(&s)
	fmt.Println(Unique(s))
}

func Unique (s string) bool {
	s=strings.ToLower(s)
	for idx:= range s{
		if strings.Count(s,string(s[idx]))>1 {
			return false
		}
	}
	return true
}