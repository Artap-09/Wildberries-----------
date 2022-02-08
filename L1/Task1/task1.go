package main

import "fmt"

func main() {
	var a Action //Создаем переменную тип Action
	fmt.Printf("%s\n%s\n",a.Go(),a.Run()) //Используем метод Go структуры Human из структуры Action
	fmt.Print(a.Now())
}

type Human struct {
	Growth uint8
	Weight uint8
	Gender string
}

func (h Human) Go() string {
	return "Иду"
}

func (h Human) Run() string {
	return "Бегу"
}

type Action struct {
	Human //встраиваем методы структуры Human аналог наследования в ОПП
	x int
	y int
	z int
}

func (a Action) Now() string {
	return fmt.Sprintf("Сейчас я нахожусь по координатам:\nx:%d y:%d z:%d\n",a.x,a.y,a.z)
}