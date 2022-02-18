package pattern

import "fmt"

/*
func main() {
	tv := &tv{}

	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
*/

// Кнопка
type button struct {
	command command
}

// Нажатие на кнопку
func (b *button) press() {
	b.command.execute()
}

//Команды
type command interface {
	execute()
}

//Команда включить
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// Команда выключить
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// Устройства
type device interface {
	on()
	off()
}

// Телевизор
type tv struct {
	isRunning bool
}

//Включить
func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

//Выключить
func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
