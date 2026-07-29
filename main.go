package main

import (
	"fmt"
)

func clearConsole() {
	fmt.Print("\033[H\033[2J")
	println()
	println()
}

type Dog struct {
	nickname    string
	age         int
	weightKg    float64
	checkInDate string
	shelter     Shelter
	policlinic  Policlinic
}

func (d Dog) printDogInfo() {

	fmt.Println("Кличка:", d.nickname)
	fmt.Println("Возраст, лет:", d.age)
	fmt.Println("Вес, кг:", d.weightKg)
	fmt.Println("Когда попал в приют:", d.checkInDate)
	fmt.Println("К какому шелтеру относится:", d.shelter.name)

}

type Shelter struct {
	name        string
	address     string
	number      string
	workingTime string
	listPets    map[string]Dog
}

func (s Shelter) printShelterInfo() {
	fmt.Println(s.name)
	fmt.Println(s.address)
	fmt.Println(s.number)
	fmt.Println(s.workingTime)
}

type Policlinic struct {
	name         string
	address      string
	number       string
	workingTime  string
	listPatients map[string]Dog
}

func inputChooseNameDog(lenSliceDog int) (intChooseNameDog int) {
	for {
		_, err := fmt.Scan(&intChooseNameDog)
		if err != nil {
			fmt.Println("Ошибка ввода")
			continue
		}
		if intChooseNameDog > 0 && intChooseNameDog <= lenSliceDog {
			break
		} else {
			fmt.Println("Ошибка диапазона")
		}

	}
	return intChooseNameDog
}

func inputPickUpDog() (intPickUpDog int) {
	for {
		_, err := fmt.Scan(&intPickUpDog)
		if err != nil {
			fmt.Println("Ошибка ввода")
			continue
		}
		if intPickUpDog != 1 && intPickUpDog != 2 {
			fmt.Println("Ошибка диапазона")
		} else {
			break
		}
	}
	return intPickUpDog
}

func main() {

	var sliceDog []Dog
	var sliceShelter []Shelter

	sliceShelter = append(sliceShelter, Shelter{
		name:        "Приют 'Шанс'",
		address:     "Пятруся Глебки 17",
		number:      "+375 29 511-22-13",
		workingTime: "10:00 - 22:00",
	})
	sliceShelter = append(sliceShelter, Shelter{
		name:        "Приют 'Привет'",
		address:     "Мстислава Чудотворца 4/1",
		number:      "+375 12 544-65-45",
		workingTime: "11:00 - 21:15",
	})

	sliceDog = append(sliceDog, Dog{
		nickname:    "Чарли",
		age:         3,
		weightKg:    7.1,
		checkInDate: "05.02.2025",
		shelter:     sliceShelter[0],
	})
	sliceDog = append(sliceDog, Dog{
		nickname:    "Вил",
		age:         4,
		weightKg:    17,
		checkInDate: "15.03.2025",
		shelter:     sliceShelter[1],
	})
	sliceDog = append(sliceDog, Dog{
		nickname:    "Кайман",
		age:         2,
		weightKg:    3,
		checkInDate: "25.11.2022",
		shelter:     sliceShelter[1],
	})
	intChooseNameDog := 0
	intPickUpDog := 0

	for {
		clearConsole()
		fmt.Println("Собаки, которых вы могли бы забрать из приюта\nЧтобы узнать больше о питомце введите его номер")
		for index, value := range sliceDog {
			fmt.Println(index+1, value.nickname)

		}

		println("Введите номер собаки, о которой хотите узнать больше:\n")
		intChooseNameDog = inputChooseNameDog((len(sliceDog))) - 1
		clearConsole()

		sliceDog[intChooseNameDog].printDogInfo()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("Забрать из приюта?\n1 Да, забрать из приюта\n2 Нет, смотреть других собак")
		intPickUpDog = inputPickUpDog()
		if intPickUpDog == 1 {
			clearConsole()
			sliceDog[intChooseNameDog].shelter.printShelterInfo()
			break
		}

	}

}
