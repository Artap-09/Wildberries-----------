package main

import "fmt"

func main() {
	roundHole:=&RoundHole{13.435}
	roundPeg1:=&RoundPeg{13.435}
	roundPeg2:=&RoundPeg{13}
	squarePeg:=&SquarePeg{19}

	squarePegAdapter:=&SquarePegAdapter{
		SquarePeg: squarePeg,
	}

	fmt.Println(roundHole.fits(roundPeg1))
	fmt.Println(roundHole.fits(roundPeg2))
	fmt.Println(roundHole.fits(squarePegAdapter))
}