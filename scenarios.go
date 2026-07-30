package main

import "fmt"

// возвращает false, если пользователь решил закончить программу
func scenarioTakeDog(dogs map[string]*Dog, shelters []Shelter, policlinics []Policlinic) bool {
	printDogList(dogs)
	nickname := readNonEmptyString("Введите кличку собаки: ")

	d, ok := FindDog(dogs, nickname)
	if !ok {
		fmt.Println("Собака с такой кличкой не найдена")
		return true // не выходим из программы, просто не нашли
	}

	printDogInfo(dogs, nickname)
	if readYesNo("Забрать из приюта? (да/нет): ") {
		RemoveDog(dogs, nickname)
		fmt.Println("Собака удалена из общего списка, приюта и поликлиники")
		printShelterInfo(d.shelter)
	}

	return readYesNo("Смотреть ещё? (да/нет): ")
}

func scenarioAddDog(dogs map[string]*Dog, shelters []Shelter, policlinics []Policlinic) bool {
	// считать nickname/age/weight/checkInDate
	// дать выбрать приют и поликлинику из списка (readMenuChoice по индексу)
	// AddDog(...) + печать сообщений на каждом шаге
	return readYesNo("Смотреть ещё? (да/нет): ")
}
