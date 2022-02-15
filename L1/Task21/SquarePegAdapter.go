package main

import "math"

type SquarePegAdapter struct {
	SquarePeg *SquarePeg
}

func (sp SquarePegAdapter) getRadius() float64{
	return math.Round((sp.SquarePeg.getWidth()*math.Sqrt2/2)*1000)/1000
}