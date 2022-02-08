package main

import (
	"fmt"
	"sync"
)

func main() {
	even := [5]int{2,4,6,8,10}
	var ans [5]int
	wg:=new(sync.WaitGroup)
	for idx, val := range even{
		wg.Add(1)//Увеличиваем счетчик горутин в группе
		go func (idx, val int)  {
			defer wg.Done()
			ans[idx]=val*val
		}(idx,val)
	}

	wg.Wait()
	fmt.Println(ans)
}