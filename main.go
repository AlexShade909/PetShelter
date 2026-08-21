package main

import (
	"fmt"
	"log"

	"PetShelter/cli"
	"PetShelter/net"
)

func main() {
	if err := net.Connect(); err != nil {
		log.Fatalln(err.Error())
	}

	shelters := CreateShelters()
	policlinics := CreatePoliclinics()
	dogs := CreateDogs(shelters, policlinics)

	flag := true

	for flag {
		choice := cli.ReadMenuChoice("1. Выбрать собаку\n2. Добавить собаку\n3. Выход\n ", 1, 3)
		switch choice {
		case 1:
			fmt.Println("Выбрать собаку, я пользователь")
			flag = scenarioTakeDog(dogs)
		case 2:
			fmt.Println("Добавить собаку, я администратор")
			flag = scenarioAddDog(dogs, shelters, policlinics)
		case 3:
			fmt.Println("Выход")
			return
		}
	}
}
