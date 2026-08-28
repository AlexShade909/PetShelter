package main

import (
	"fmt"
	//"PetShelter/cli"
)

func scenarioTakeDog(dogs map[string]Dog) bool {
	printDogList(dogs)
	nickname := ReadNonEmptyString("Введите кличку собаки: ")
	d, ok := FindDog(dogs, nickname)
	if !ok {
		fmt.Println("Собака с такой кличкой не найдена")
		return true
	}
	printDogInfo(dogs, nickname)
	if ReadYesNo("Забрать из приюта? (да/нет): ") {
		RemoveDog(dogs, nickname)
		fmt.Println("Собака удалена из общего списка, приюта и поликлиники")
		printShelterInfo(d.shelter)
		printPoliclinicInfo(d.policlinic)
	}

	return ReadYesNo("Смотреть ещё? (да/нет): ")
}

func scenarioAddDog(dogs map[string]Dog, shelters []Shelter, policlinics []Policlinic) bool {
	nickname := ReadNonEmptyString("Введите кличку: ")
	age := ReadInt("Введите возраст: ")
	weight := ReadFloat("Введите вес: ")
	date := ReadNonEmptyString("Введите дату поступления: ")
	shelterChoice := ReadMenuChoice("Выберите приют: ", 0, len(shelters)-1)
	shelter := &shelters[shelterChoice]
	clinicChoice := ReadMenuChoice("Выберите поликлинику: ", 0, len(policlinics)-1)
	policlinic := &policlinics[clinicChoice]
	dog := AddDog(
		dogs,
		nickname,
		age,
		weight,
		date,
		shelter,
		policlinic,
	)
	fmt.Println("Собака добавлена:", dog.nickname)

	return ReadYesNo("Смотреть ещё? (да/нет): ")
}
