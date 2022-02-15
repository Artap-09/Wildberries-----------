package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int
	fmt.Println("Введите количество воркеров:")
	fmt.Scanln(&count)
	ch:= make(chan interface{})
	for i := 0; i < count; i++ {
		go worker(ch,i+1)
	}

	for{
		rand.Seed(time.Now().UnixNano())
		ch<-rand.Intn(101)
		ch<-rand.Float64()
		ch<-string(byte(rand.Intn(26)+65))
	}
}

func worker(ch <-chan interface{},idx int)  {
	for{
		fmt.Printf("Я воркер №%d: %v\n",idx,<-ch)
		time.Sleep(time.Second)
	}
}