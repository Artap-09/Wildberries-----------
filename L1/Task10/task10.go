package main

import (
	"fmt"
	"sync"
)

func main() {
	temperature:=make(map[int][]float64)
	wg:= new(sync.WaitGroup)
	ch:=make(chan float64)

	wg.Add(2)
	go Sender(ch, wg)
	go Worker(ch, temperature, wg)

	wg.Wait()
	fmt.Println(temperature)
}

func Sender(chanOut chan<- float64, wg *sync.WaitGroup)  {
	defer wg.Done()
	m:= [8]float64 {-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, v := range m {
		chanOut<-v
	}
	
	close(chanOut)
}

func Worker(chanIn <-chan float64, m map[int][]float64, wg *sync.WaitGroup)  {
	defer wg.Done()
	for v := range chanIn {
		m[(int(v)/10)*10]=append(m[(int(v)/10)*10], v)
	}
}