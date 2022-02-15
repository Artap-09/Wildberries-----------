package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	Sleep(time.Second)
	fmt.Println(time.Now())
}

func Sleep(d time.Duration)  {
	if d<=0 {
		return
	}
	
	<-time.After(d)
}