package main

import "fmt"

func printDogInfo(dogs map[string]*Dog, nickname string) {
	fmt.Println("Кличка:", nickname)
	fmt.Println("Возраст, лет:", dogs[nickname].age)
	fmt.Println("Вес, кг:", dogs[nickname].weightKg)
	fmt.Println("Когда попал в приют:", dogs[nickname].checkInDate)
	fmt.Println("К какому шелтеру относится:", dogs[nickname].shelter.address)
	fmt.Println("К какой поликлинике относится:", dogs[nickname].policlinic.address)
}

func printDogList(dogs map[string]*Dog) {
	fmt.Println("Список собак: ")
	for nickname := range dogs {
		fmt.Println("-", nickname)
	}
}

func printShelterInfo(s *Shelter) {
	fmt.Println("Приют №:", s.numberShelter)
	fmt.Println("Адрес:", s.address)
	fmt.Println("Телефон:", s.number)
	fmt.Println("Время работы:", s.workingTime)
	fmt.Println("Собак в приюте:", len(s.listPets))
}
