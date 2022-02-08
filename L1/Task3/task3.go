package main

import (
	"fmt"
	"sync"
)

func main() {
	even := [5]int{2,4,6,8,10}
	var ans int
	wg:=new(sync.WaitGroup)
	for _, val := range even{
		wg.Add(1)//Увеличиваем счетчик горутин в группе
		go func ( val int)  {
			defer wg.Done()
			ans+=val*val
		}(val)
	}

	wg.Wait()
	fmt.Println(ans)
}