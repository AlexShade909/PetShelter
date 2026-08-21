package main

import "fmt"

func main() {
	shelters := CreateShelters()
	policlinics := CreatePoliclinics()
	dogs := CreateDogs(shelters, policlinics)
	fmt.Println(dogs)
	/*
		facadeStorage := FacadeStorage{
			PoliclinicsStorage: policlinicsStorage,
			SheltersStorage:    sheltersStorage,
		}

		flag := true

			for flag {
				choice := readMenuChoice("1. Выбрать собаку\n2. Добавить собаку\n3. Выход\n ", 1, 3)
				switch choice {
				case 1:
					fmt.Println("Выбрать собаку")
					flag = scenarioTakeDog(facade)
				case 2:
					fmt.Println("Добавить собаку")
					flag = scenarioAddDog(facade)
				case 3:
					fmt.Println("Выход")
					return
				default:
					fmt.Println("неизвестная команда")
					//такой вариант исключён логикой функции readMenuChoice
				}
			}

	*/
}
