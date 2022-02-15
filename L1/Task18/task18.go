package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg:=new(sync.WaitGroup)
	count:=NewCount()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(count,i,wg)
	}

	wg.Wait()
	fmt.Printf("Было выполнено %d горутин\n",count.count)
}

type Count struct{
	count int
	sync.RWMutex
}

func NewCount()*Count{
	return &Count{}
}

func (c *Count) Increase(){
	c.Lock()
	c.count++
	c.Unlock()
}

func worker(c *Count,i int, wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Printf("Горутина %d начала выполнение\n", i)
	time.Sleep(2*time.Second)
	fmt.Printf("Горутина %d завершила выполнение\n", i)
	c.Increase()
}