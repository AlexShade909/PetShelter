package main

import "fmt"

type Dog struct {
	nickname    string
	age         int
	weightKg    float64
	checkInDate string
	shelter     *Shelter
	policlinic  *Policlinic
}

func (d *Dog) Print() {
	fmt.Println("Кличка:", d.nickname)
	fmt.Println("Возраст, лет:", d.age)
	fmt.Println("Вес, кг:", d.weightKg)
	fmt.Println("Когда попал в приют:", d.checkInDate)
	fmt.Println("К какому шелтеру относится:", d.shelter.address)
	fmt.Println("К какой поликлинике относится:", d.policlinic.address)
}

type Shelter struct {
	numberShelter string
	address       string
	number        string
	workingTime   string
}

type Policlinic struct {
	numberClinic string
	address      string
	phoneNumber  string
	workingTime  string
}
