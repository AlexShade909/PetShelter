package main

import "fmt"

func main() {

	policlinics := CreatePoliclinics()
	shelters := CreateShelters()
	dogs := CreateDogs(shelters, policlinics)

	choice := readMenuChoice("1. Выбрать собаку\n2. Добавить собаку\n3. Выход\n ", 1, 3)
	switch choice {
	case 1:
		fmt.Println("Выбрать собаку")
		for scenarioTakeDog(dogs, shelters, policlinics) {
		}
	case 2:
		fmt.Println("Добавить собаку")
		scenarioAddDog(dogs, shelters, policlinics)
		return

	case 3:
		fmt.Println("Выход")
		return

	default:
		fmt.Println("неизвестная команда")
		//такой вариант исключён логикой функции readMenuChoice
	}

}
