package main

import (
	"fmt"
	"math/big"
)

func main() {
	x:=big.NewInt(56e17)
	y:=big.NewInt(28e16)
	result:=new(big.Int)

	fmt.Println(result.Div(x,y))
	fmt.Println(result.Add(x,y))
	fmt.Println(result.Mul(x,y))
	fmt.Println(result.Sub(x,y))
}