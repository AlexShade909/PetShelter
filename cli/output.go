package cli

import (
	"PetShelter/internal"
	"fmt"
)

func PrintDogInfo(dogs map[string]internal.Dog, nickname string) {
	d := dogs[nickname]
	fmt.Println("Кличка:", d.Nickname)
	fmt.Println("Возраст, лет:", d.Age)
	fmt.Println("Вес, кг:", d.WeightKg)
	fmt.Println("Когда попал в приют:", d.CheckInDate)
	fmt.Println("К какому шелтеру относится:", d.Shelter.Address)
	fmt.Println("К какой поликлинике относится:", d.Policlinic.Address)
}

func PrintDogList(dogs map[string]internal.Dog) {
	fmt.Println("Список собак: ")
	for nickname := range dogs {
		fmt.Println(nickname)
	}
}

func PrintShelterInfo(s *internal.Shelter) {
	fmt.Println("Приют №:", s.NumberShelter)
	fmt.Println("Адрес:", s.Address)
	fmt.Println("Телефон:", s.Number)
	fmt.Println("Время работы:", s.WorkingTime)
}

func PrintPoliclinicInfo(p *internal.Policlinic) {
	fmt.Println("Название:", p.NumberClinic)
	fmt.Println("адрес:", p.Address)
	fmt.Println("телефон:", p.PhoneNumber)
	fmt.Println("рабочее время:", p.WorkingTime)
}
