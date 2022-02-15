package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg:=new(sync.WaitGroup)
	quit:=make(chan struct{})
	ch:= make(chan string)
	ch1:= make(chan string)

	wg.Add(1)
	go func() {
		for {
			select{
			case <-quit: // Закрытие произайдет при отправки пустой структуры (иногда используют bool)
				fmt.Println("Закрыли горутину №1")
				wg.Done()
				return
			case <-ch1:
				fmt.Println("Я горутина №1")
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Add(1)
	go func() {
		for {
			if _, ok:=<-ch1; !ok{ // Закрытие произайдет при закрытие канала
				fmt.Println("Закрыли горутину №2")
				wg.Done()
				return
			}
			fmt.Println("Я горутина №2")
			time.Sleep(time.Second)

		}
	}()

	wg.Add(1)
	go func() {
		for {
			switch <-ch {
			case "close": // Закрытие произайдет при отправки в канал тригера (в данном случае строки "close")
				fmt.Println("Закрыли горутину №3")
				wg.Done()
				return
			default:
				fmt.Println("Я горутина №3")
				time.Sleep(time.Second)
			}				
		}
	}()

	i:=0
	wg.Add(1)
	go func() {
		for{
			i++
			select {
			case <-time.After(2*time.Second): // закрывает канал через 2 секунды если select не провалится во второй case 
				fmt.Println("Закрыли горутину №4")
				wg.Done()
				return
			case<-ch:
				fmt.Println("Я горутина №4")
				time.Sleep(time.Second)
			}
		}
	}()

	for  i!=5 {
		ch<-""	
		ch1<-""	
	}

	quit<-struct{}{}
	ch<-"close"
	close(ch1)
	wg.Wait()
	fmt.Println("Все горутины закрыты")
}