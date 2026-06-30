package main

import (
	"fmt"
)

func clearConsole() {
	fmt.Print("\033[H\033[2J")
	println()
	println()
}

type Shelter struct {
	nameShelter   string
	address       string
	contactNumber string
	workingTime   string
}

type Dog struct {
	numberInBase   int
	nickname       string
	age            int
	weightKg       int
	checkInDate    string
	stayShelerName string
}

func (d Dog) printDogInfo() {

	println("Кличка:", d.nickname)
	println("Номер в базе:", d.numberInBase)
	println("Возраст, лет:", d.age)
	println("Вес, кг:", d.weightKg)
	println("Когда попал в приют:", d.checkInDate)
	println("К какому шелтеру относится:", d.stayShelerName)

}

func (d Shelter) printShelterInfo() {
	println(d.nameShelter)
	println(d.address)
	println(d.contactNumber)
	println(d.workingTime)
}

func main() {

	Shelter1 := Shelter{
		nameShelter:   "Приют 'Шанс'",
		address:       "Пятруся Глебки 17",
		contactNumber: "+375 29 511-22-13",
		workingTime:   "10:00 - 22:00",
	}

	Shelter2 := Shelter{
		nameShelter:   "Приют 'Привет'",
		address:       "Мстислава Чудотворца 4/1",
		contactNumber: "+375 12 544-65-45",
		workingTime:   "11:00 - 21:15",
	}

	Dog1 := Dog{
		numberInBase:   1,
		nickname:       "Чарли",
		age:            3,
		weightKg:       7,
		checkInDate:    "05.02.2025",
		stayShelerName: Shelter1.nameShelter,
	}

	Dog2 := Dog{
		numberInBase:   2,
		nickname:       "Вил",
		age:            4,
		weightKg:       17,
		checkInDate:    "15.03.2025",
		stayShelerName: Shelter2.nameShelter,
	}
	Dog3 := Dog{
		numberInBase:   3,
		nickname:       "Кайман",
		age:            2,
		weightKg:       3,
		checkInDate:    "25.11.2022",
		stayShelerName: Shelter1.nameShelter,
	}

	for {
		varChooseNameDog := 0
		varPickUpDog := 0

		clearConsole()

		println("Собаки, которых вы могли бы забрать из приюта\nЧтобы узнать больше о питомце введите его номер")
		println(Dog1.numberInBase, Dog1.nickname)
		println(Dog2.numberInBase, Dog2.nickname)
		println(Dog3.numberInBase, Dog3.nickname)

		println("Введите номер собаки, о которой хотите узнать больше:\n")
		fmt.Scan(&varChooseNameDog)

		clearConsole()

		if varChooseNameDog == Dog1.numberInBase {
			Dog1.printDogInfo()

			println()
			println()
			println()
			println("Забрать из приюта?\n1 Да, забрать из приюта\n2 Нет, смотреть других собак")
			fmt.Scan(&varPickUpDog)

			if varPickUpDog == 1 {
				clearConsole()

				Shelter1.printShelterInfo()

				break

			}

		}
		if varChooseNameDog == Dog2.numberInBase {
			Dog2.printDogInfo()

			println()
			println()
			println()
			println("Забрать из приюта?\n1 Да, забрать из приюта\n2 Нет, смотреть других собак")
			fmt.Scan(&varPickUpDog)

			if varPickUpDog == 1 {
				clearConsole()
				Shelter2.printShelterInfo()
				break

			}
		}
		if varChooseNameDog == Dog3.numberInBase {
			Dog3.printDogInfo()

			println()
			println()
			println()
			println("Забрать из приюта?\n1 Да, забрать из приюта\n2 Нет, смотреть других собак")
			fmt.Scan(&varPickUpDog)

			if varPickUpDog == 1 {
				clearConsole()
				Shelter1.printShelterInfo()
				break

			}
		}

	}

}
