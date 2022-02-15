package main

import "fmt"

func main() {
	str:=[]string{"A","B","C","D","E"}
	for str!=nil {
		str=deleteEl(0,str)
		fmt.Println(str)
	}
}

func deleteEl(i int,s []string) []string {
	if len(s)==0 {
		return nil
	}

	s[i]=s[len(s)-1]
	s=s[:len(s)-1]
	return s
}