package main

import (
	"fmt"
	"math"
)

func main() {
	p:=&point{5,7}
	p1:=&point{2,3}

	fmt.Println(p.constructor(p1))
}

type point struct{
	x float64
	y float64
}

func (p *point) constructor(p2 *point) float64{
	return math.Sqrt(math.Pow(p.x-p2.x,2)+math.Pow(p.y-p2.y,2))
}