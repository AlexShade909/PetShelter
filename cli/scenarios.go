package cli

import (
	"PetShelter/internal"
	"fmt"
)

func ScenarioTakeDog(dogs map[string]internal.Dog) bool {
	PrintDogList(dogs)
	nickname := ReadNonEmptyString("Введите кличку собаки: ")
	d, ok := internal.FindDog(dogs, nickname)
	if !ok {
		fmt.Println("Собака с такой кличкой не найдена")
		return true
	}
	PrintDogInfo(dogs, nickname)
	if ReadYesNo("Забрать из приюта? (да/нет): ") {
		internal.RemoveDog(dogs, nickname)
		fmt.Println("Собака удалена из общего списка, приюта и поликлиники")
		PrintShelterInfo(d.Shelter)
		PrintPoliclinicInfo(d.Policlinic)
	}

	return ReadYesNo("Смотреть ещё? (да/нет): ")
}

func ScenarioAddDog(dogs map[string]internal.Dog, shelters []internal.Shelter, policlinics []internal.Policlinic) bool {
	nickname := ReadNonEmptyString("Введите кличку: ")
	age := ReadInt("Введите возраст: ")
	weight := ReadFloat("Введите вес: ")
	date := ReadNonEmptyString("Введите дату поступления: ")
	shelterChoice := ReadMenuChoice("Выберите приют: ", 0, len(shelters)-1)
	shelter := &shelters[shelterChoice]
	clinicChoice := ReadMenuChoice("Выберите поликлинику: ", 0, len(policlinics)-1)
	policlinic := &policlinics[clinicChoice]
	dog := internal.AddDog(
		dogs,
		nickname,
		age,
		weight,
		date,
		shelter,
		policlinic,
	)
	fmt.Println("Собака добавлена:", dog.Nickname)

	return ReadYesNo("Смотреть ещё? (да/нет): ")
}
