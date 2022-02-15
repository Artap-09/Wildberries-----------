package main

type RoundPeg struct{
	radius float64
}

func (rp RoundPeg) getRadius() float64  {
	return rp.radius
}