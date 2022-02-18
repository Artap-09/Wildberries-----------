package pattern

import "fmt"

/*
func main() {

    cashier := &cashier{}

    //Set next for medical department
    medical := &medical{}
    medical.setNext(cashier)

    //Set next for doctor department
    doctor := &doctor{}
    doctor.setNext(medical)

    //Set next for reception department
    reception := &reception{}
    reception.setNext(doctor)

    patient := &patient{name: "abc"}
    //Patient visiting
    reception.execute(patient)
}
*/

// Отделение
type department interface {
	execute(*patient)
	setNext(department)
}

// Ресепшен
type reception struct {
	next department
}

// Выполнить
func (r *reception) execute(p *patient) {
	if p.registrationDone { // Если пациент уже был в регистратуре
		fmt.Println("Patient registration already done")
		r.next.execute(p) // Переходим к следующему этапу
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p) // Переходим к следующему этапу
}

// Переход к след. этапу
func (r *reception) setNext(next department) {
	r.next = next
}

// Доктор
type doctor struct {
	next department
}

// Выполнить
func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone { // Если пациент уже был у пациента
		fmt.Println("Doctor checkup already done")
		d.next.execute(p) // Переходим к следующему этапу
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p) // Переходим к следующему этапу
}

// Переход к след. этапу
func (d *doctor) setNext(next department) {
	d.next = next
}

// Медикоменты
type medical struct {
	next department
}

// Выполнить
func (m *medical) execute(p *patient) {
	if p.medicineDone { // Если уже получили медикоменты
		fmt.Println("Medicine already given to patient")
		m.next.execute(p) // Переходим к следующему этапу
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p) // Переходим к следующему этапу
}

// Переход к след. этапу
func (m *medical) setNext(next department) {
	m.next = next
}

// Касса
type cashier struct {
	next department
}

// Выполнить
func (c *cashier) execute(p *patient) {
	if p.paymentDone { // Если уже оплачено
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next department) { // Переходим к следующему этапу
	c.next = next
}

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}
