package cli

import (
	"petshelter/internal"
)

func ScenarioTakeDog(dogs map[string]internal.Dog) bool {
	PrintDogList(dogs)
	nickname := ValidationReadNonEmptyString("Введите кличку собаки: ")
	d, ok := internal.FindDog(dogs, nickname)
	if !ok {
		Println("Собака с такой кличкой не найдена")
		return true
	}
	PrintDogInfo(dogs, nickname)
	if ValidationReadYesNo("Забрать из приюта? (да/нет): ") {
		internal.RemoveDog(dogs, nickname)
		Println("Собака удалена из общего списка, приюта и поликлиники")
		PrintShelterInfo(d.Shelter)
		PrintPoliclinicInfo(d.Policlinic)
	}

	return ValidationReadYesNo("Смотреть ещё? (да/нет): ")
}

func ScenarioAddDog(dogs map[string]internal.Dog, shelters []internal.Shelter, policlinics []internal.Policlinic) bool {
	nickname := ValidationReadNonEmptyString("Введите кличку: ")
	age := ValidationReadNonEmptyString("Введите возраст: ")
	weight := ValidationReadNonEmptyString("Введите вес: ")
	date := ValidationReadNonEmptyString("Введите дату поступления: ")
	shelterChoice := ValidationReadMenuChoice("Выберите приют: ", 0, len(shelters)-1)
	shelter := &shelters[shelterChoice]
	clinicChoice := ValidationReadMenuChoice("Выберите поликлинику: ", 0, len(policlinics)-1)
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
	Println("Собака добавлена:")
	Println(dog.Nickname)

	return ValidationReadYesNo("Смотреть ещё? (да/нет): ")
}
