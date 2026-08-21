package main

import "fmt"

func main() {
	policlinicsStorage := PoliclinicsStorage{
		table: make(map[string]Policlinic),
	}
	policlinicsStorage.Add(CreatePoliclinics()...)

	sheltersStorage := SheltersStorage{
		table: make(map[string]Shelter),
	}
	sheltersStorage.Add(CreateShelters()...)

	policlinnic1 := policlinicsStorage.Find("Поликлиника 1")
	policlinnic2 := policlinicsStorage.Find("Поликлиника 2")
	shelter1 := sheltersStorage.Find("Шелтер 1")
	shelter2 := sheltersStorage.Find("Шелтер 2")

	fmt.Println(policlinnic1)
	fmt.Println(policlinnic2)
	fmt.Println(shelter1)
	fmt.Println(shelter2)
	fmt.Println()
	fmt.Println()

	dogs := CreateDogs()
	fmt.Println(dogs)
	fmt.Println()
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
