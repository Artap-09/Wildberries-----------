package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	mu:=new(sync.Mutex)
	myMap:=make(map[string]string)
	done:=make(chan struct{})
	chKeyValue:=make(chan string)

	nWriter:=10

	for i := 0; i < nWriter; i++ {
		go Writer(chKeyValue, done, mu, myMap)
	}

	for i := 0; i < nWriter; i++ {
		rand.Seed(time.Now().UnixNano())
		chKeyValue<- string(byte(rand.Intn(26)+65))+": "+string(byte(rand.Intn(26)+65))
	}

	for i := 0; i < nWriter; i++ {
		<-done
	}

	fmt.Println(myMap)

}

func Writer(chKeyValue <-chan string,done chan<- struct{}, mu *sync.Mutex,myMap map[string]string)  {
	strs:=strings.Split(<-chKeyValue,": ")

	if len(strs)!=2 {
		log.Println("Некорректный формат")
		done<-struct{}{}
		return
	}
	mu.Lock()
	myMap[strs[0]]=strs[1]
	fmt.Printf("[%s] Ключ: %s Значение: %s\n",time.Now().Format("2006-01-02 15:04:05.000000"),strs[0],myMap[strs[0]])
	mu.Unlock()

	done<-struct{}{}
}