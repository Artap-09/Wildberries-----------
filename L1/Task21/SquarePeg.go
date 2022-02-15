package main
type SquarePeg struct{
	width float64
}

func (sp SquarePeg) getWidth() float64 {
	return sp.width
}