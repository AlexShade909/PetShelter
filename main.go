package main

import (
	"PetShelter/cli"
	"fmt"
)

func main() {
	//if err := net.Connect(); err != nil {
	//	log.Fatalln(err.Error())
	//
	Shelters := CreateShelters()
	Policlinics := CreatePoliclinics()
	Dogs := CreateDogs(Shelters, Policlinics)

	flag := true

	for flag {
		choice := cli.ReadMenuChoice("1. Выбрать собаку\n2. Добавить собаку\n3. Выход\n ", 1, 3)
		switch choice {
		case 1:
			fmt.Println("Выбрать собаку, я пользователь")
			flag = scenarioTakeDog(Dogs)
		case 2:
			fmt.Println("Добавить собаку, я администратор")
			flag = scenarioAddDog(Dogs, Shelters, Policlinics)
		case 3:
			fmt.Println("Выход")
			return
		}
	}
}
