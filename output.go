package main

import "fmt"

func printDogInfo(dogs map[string]Dog, nickname string) {
	dog := dogs[nickname]

	dog.Print()
}

func printDogList(dogs map[string]Dog) {
	fmt.Println("Список собак: ")
	for nickname := range dogs {
		fmt.Println(nickname)
	}
}

func printShelterInfo(s *Shelter) {
	fmt.Println("Приют №:", s.numberShelter)
	fmt.Println("Адрес:", s.address)
	fmt.Println("Телефон:", s.number)
	fmt.Println("Время работы:", s.workingTime)
}

func printPoliclinicInfo(p *Policlinic) {
	fmt.Println("Название:", p.numberClinic)
	fmt.Println("адрес:", p.address)
	fmt.Println("телефон:", p.phoneNumber)
	fmt.Println("рабочее время:", p.workingTime)
}
