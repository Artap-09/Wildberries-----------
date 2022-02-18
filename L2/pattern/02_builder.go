package pattern

/*
import "fmt"


func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}
*/

// Интерфейс строителя
type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

// Возвращает строителя который необходим
func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

// Строитель 1
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// Создает строителя 1
func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

// Устанавливает окно
func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

// Устанавливает дверь
func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

// Устанавливает количество этажей
func (b *normalBuilder) setNumFloor() {
	b.floor = 2
}

// Возвращает готовый дом
func (b *normalBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// Строитель 2
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// Создает строителя 2
func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

// Устанавливает окно
func (b *iglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

// Устанавливает дверь
func (b *iglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

// Устанавливает количество этажей
func (b *iglooBuilder) setNumFloor() {
	b.floor = 1
}

// Возвращает готовый дом
func (b *iglooBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// Продукт
type house struct {
	windowType string
	doorType   string
	floor      int
}

// Директор
type director struct {
	builder iBuilder
}

// Создает директора
func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

// Назначает строителя
func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

// Строит дом
func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
