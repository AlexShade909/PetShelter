package main

import "fmt"

// возвращает false, если пользователь решил закончить программу
func scenarioTakeDog(facade Facade) bool {
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

func scenarioAddDog(facade Facade) bool {
	// считать nickname/age/weight/checkInDate
	// дать выбрать приют и поликлинику из списка (readMenuChoice по индексу)
	// AddDog(...) + печать сообщений на каждом шаге
	nickname := readNonEmptyString("Введите кличку: ")
	age := readInt("Введите возраст: ")
	weight := readFloat("Введите вес: ")
	date := readNonEmptyString("Введите дату поступления: ")
	// выбрать приют
	shelterChoice := readMenuChoice("Выберите приют: ", 1, len(shelters))
	shelter := &shelters[shelterChoice-1]
	// выбрать поликлинику
	clinicChoice := readMenuChoice("Выберите поликлинику: ", 1, len(policlinics))
	policlinic := &policlinics[clinicChoice-1]
	dog := AddDog(
		dogs,
		nickname,
		age,
		weight,
		date,
		shelter,
		policlinic,
	)
	fmt.Println("Собака добавлена:", dog.nickname)
	return readYesNo("Смотреть ещё? (да/нет): ")
}
