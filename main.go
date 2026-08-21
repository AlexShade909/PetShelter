package main

import "fmt"

func main() {
	policlinicsStorage := PoliclinicsStorage{}
	policlinicsStorage.Add(CreatePoliclinics()...)

	sheltersStorage := SheltersStorage{}
	sheltersStorage.Add(CreateShelters()...)

	fmt.Println(sheltersStorage.Find("1"))
	fmt.Println("Remove")
	sheltersStorage.Remove("1")
	fmt.Println(sheltersStorage.Find("1"))

	/*

		dogs := CreateDogs()

		policlinnic1 := policlinicsStorage.Find("1")
		policlinnic2 := policlinicsStorage.Find("2")
		shelter1 := sheltersStorage.Find("1")
		shelter2 := sheltersStorage.Find("2")

		policlinnic1.AddDogs(dogs[0], dogs[1])
		policlinnic2.AddDogs(dogs[2])

		shelter1.AddDogs(dogs[0], dogs[2])
		shelter2.AddDogs(dogs[1])


			facade := FacadeStorage{
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
