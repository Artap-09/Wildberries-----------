package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var n int
	wg:=new(sync.WaitGroup)
	ch := make(chan int)
	quit := make(chan int)
	fmt.Scanln(&n)
	d,_:=time.ParseDuration(fmt.Sprint(n)+"s")
	timer:=time.AfterFunc(d, func() {
		quit<-0
		close(quit)
	})
	defer timer.Stop()

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		i:=1
		for{			
			select{
			case ch <- i:
				i++
			case <-quit:
				return
			}
		}		
	} ()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select{
			case <-ch:
			case <-quit:
				return
			}
		}
	}()

	wg.Wait()
}	