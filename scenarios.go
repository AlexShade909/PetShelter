package main

import "fmt"

func scenarioTakeDog(dogs map[string]Dog) bool {
	printDogList(dogs)
	nickname := readNonEmptyString("Введите кличку собаки: ")
	d, ok := FindDog(dogs, nickname)
	if !ok {
		fmt.Println("Собака с такой кличкой не найдена")
		return true
	}
	printDogInfo(dogs, nickname)
	if readYesNo("Забрать из приюта? (да/нет): ") {
		RemoveDog(dogs, nickname)
		fmt.Println("Собака удалена из общего списка, приюта и поликлиники")
		printShelterInfo(d.shelter)
		printPoliclinicInfo(d.policlinic)
	}

	return readYesNo("Смотреть ещё? (да/нет): ")

}

func scenarioAddDog(dogs map[string]Dog, shelters []Shelter, policlinics []Policlinic) bool {
	nickname := readNonEmptyString("Введите кличку: ")
	age := readInt("Введите возраст: ")
	weight := readFloat("Введите вес: ")
	date := readNonEmptyString("Введите дату поступления: ")
	shelterChoice := readMenuChoice("Выберите приют: ", 0, len(shelters)-1)
	shelter := &shelters[shelterChoice]
	clinicChoice := readMenuChoice("Выберите поликлинику: ", 0, len(policlinics)-1)
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

	return readYesNo("Смотреть ещё? (да/нет): ")
}
