package main

import (
	"fmt"
	"sync"
)

func main() {
	chIn:=make(chan int)
	chOut:=make(chan int)

	wg:= new(sync.WaitGroup)

	wg.Add(3)

	go Sender(chIn,wg)
	go Pow(chIn,chOut,wg)
	go Printer(chOut,wg)

	wg.Wait()
}

func Pow(chIn <-chan int, chOut chan<- int,wg *sync.WaitGroup)  {
	defer wg.Done()

	for v := range chIn {
		chOut <- v*v
	}

	close(chOut)
}

func Sender(chOut chan<- int,wg *sync.WaitGroup) {
	defer wg.Done()

	for _, v := range [10]int{31,142,23,3,12,87,32,56,23,54} {
		chOut<- v
	}

	close(chOut)
}

func Printer(chIn <-chan int,wg *sync.WaitGroup) {
	defer wg.Done()

	for v:= range chIn {
		fmt.Println(v)
	}
}