package main

import (
	"PetShelter/cli"
	"PetShelter/internal"
	"PetShelter/network"
	"fmt"
	"log"
)

func main() {
	if err := network.Connect(); err != nil {
		log.Fatalln(err.Error())
	}

	Shelters := internal.CreateShelters()
	Policlinics := internal.CreatePoliclinics()
	Dogs := internal.CreateDogs(Shelters, Policlinics)

	flag := true

	for flag {
		choice := cli.ReadMenuChoice("1. Выбрать собаку\n2. Добавить собаку\n3. Выход\n ", 1, 3)
		switch choice {
		case 1:
			fmt.Println("Выбрать собаку, я пользователь")
			flag = cli.ScenarioTakeDog(Dogs)
		case 2:
			fmt.Println("Добавить собаку, я администратор")
			flag = cli.ScenarioAddDog(Dogs, Shelters, Policlinics)
		case 3:
			fmt.Println("Выход")
			return
		}
	}
}
