package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan []int)
	done := make(chan struct{})
	m:= [10] int {1,8,9,5,13,15,3,6,4,2}
	wg:=new(sync.WaitGroup)

	wg.Add(2)
	go Sender(ch1,wg)
	go Reader(ch2,ch3,wg)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go Worker(ch1, ch2,done,m,wg)
	}
	
	for i := 0; i < 4; i++ {
		<-done
	} 

	close(ch2)
	close(done)

	fmt.Println(<-ch3)

	wg.Wait()
	
}

func Sender(chOut chan<- int,wg *sync.WaitGroup)  {
	defer wg.Done()
	for _, v := range [7]int{1,2,3,4,5,6,7} {
		chOut<-v
	}
	close(chOut)
}

func Worker(chIn <-chan int, chOut chan<- int,done chan<- struct{}, m [10]int,wg *sync.WaitGroup){
	defer wg.Done()
	for v := range chIn {
		for _, val := range m {
			if v==val {
				chOut<-v
				break
			}
		}
	}
	done<-struct{}{}
}

func Reader(chIn <-chan int,chOut chan<- []int, wg *sync.WaitGroup)  {
	defer wg.Done()
	s:=make([]int,0,7)
	for v:= range chIn{
		s=append(s,v)
	}

	chOut<-s
	close(chOut)
}