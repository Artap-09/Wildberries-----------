package main

type RoundHole struct {
	radius float64
}

func (rh RoundHole) getRadius() float64 {
	return rh.radius
}

func (rh RoundHole) fits(p Peg) bool {
	return rh.getRadius() == p.getRadius()
}
