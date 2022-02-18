package pattern

import (
	"fmt"
	"math"
)

/*
func main() {
	square := &square{side: 2}
	circle := &circle{radius: 3}
	rectangle := &rectangle{l: 2, b: 3}

	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &middleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
*/

// Посетитель
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForrectangle(*rectangle)
}

// Вычислить площадь
type areaCalculator struct {
	area float64
}

func (a *areaCalculator) visitForSquare(s *square) {
	a.area = s.side * s.side
	fmt.Printf("Calculating area for square. %.2f\n", a.area)
}

func (a *areaCalculator) visitForCircle(s *circle) {
	a.area = s.radius * s.radius * math.Pi
	fmt.Printf("Calculating area for circle. %.2f\n", a.area)
}
func (a *areaCalculator) visitForrectangle(s *rectangle) {
	a.area = s.b * s.l
	fmt.Printf("Calculating area for rectangle. %.2f\n", a.area)
}

// Расчитать координаты средней точки
type middleCoordinates struct {
	x float64
	y float64
}

func (a *middleCoordinates) visitForSquare(s *square) {
	a.x=s.side/2
	a.y=s.side/2
	fmt.Printf("Calculating middle point coordinates for square. %.2f %.2f\n", a.x,a.y)
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	a.x=c.radius/math.Sqrt2
	a.y=c.radius/math.Sqrt2
	fmt.Printf("Calculating middle point coordinates for circle %.2f %.2f\n", a.x,a.y)
}

func (a *middleCoordinates) visitForrectangle(t *rectangle) {
	a.x=t.b/2
	a.y=t.l/2
	fmt.Printf("Calculating middle point coordinates for rectangle %.2f %.2f\n", a.x,a.y)
}

// Элемент
type shape interface {
	getType() string
	accept(visitor)
}

// Квадрат
type square struct {
	side float64
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

// Круг
type circle struct {
	radius float64
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

// Прямоугольник
type rectangle struct {
	l float64
	b float64
}

func (t *rectangle) accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
	return "rectangle"
}