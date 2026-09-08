package models

import "fmt"

type Dog struct {
	Nickname    string
	Age         string
	WeightKg    string
	CheckInDate string
	Shelter     *Shelter
	Policlinic  *Policlinic
}

func (d *Dog) Print() {
	fmt.Println("Кличка:", d.Nickname)
	fmt.Println("Возраст, лет:", d.Age)
	fmt.Println("Вес, кг:", d.WeightKg)
	fmt.Println("Когда попал в приют:", d.CheckInDate)
	fmt.Println("К какому шелтеру относится:", d.Shelter.Address)
	fmt.Println("К какой поликлинике относится:", d.Policlinic.Address)
}
