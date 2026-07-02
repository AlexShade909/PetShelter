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
	nickname       string
	age            int
	weightKg       int
	checkInDate    string
	stayShelerName Shelter
}

func (d Dog) printDogInfo() {

	fmt.Println("Кличка:", d.nickname)
	fmt.Println("Возраст, лет:", d.age)
	fmt.Println("Вес, кг:", d.weightKg)
	fmt.Println("Когда попал в приют:", d.checkInDate)
	fmt.Println("К какому шелтеру относится:", d.stayShelerName.nameShelter)

}

func (d Shelter) printShelterInfo() {
	fmt.Println(d.nameShelter)
	fmt.Println(d.address)
	fmt.Println(d.contactNumber)
	fmt.Println(d.workingTime)
}

func main() {

	var sliceDog []Dog
	var sliceShelter []Shelter

	sliceShelter = append(sliceShelter, Shelter{
		nameShelter:   "Приют 'Шанс'",
		address:       "Пятруся Глебки 17",
		contactNumber: "+375 29 511-22-13",
		workingTime:   "10:00 - 22:00",
	})

	sliceShelter = append(sliceShelter, Shelter{
		nameShelter:   "Приют 'Привет'",
		address:       "Мстислава Чудотворца 4/1",
		contactNumber: "+375 12 544-65-45",
		workingTime:   "11:00 - 21:15",
	})

	sliceDog = append(sliceDog, Dog{
		nickname:       "Чарли",
		age:            3,
		weightKg:       7,
		checkInDate:    "05.02.2025",
		stayShelerName: sliceShelter[0],
	})

	sliceDog = append(sliceDog, Dog{
		nickname:       "Вил",
		age:            4,
		weightKg:       17,
		checkInDate:    "15.03.2025",
		stayShelerName: sliceShelter[1],
	})
	sliceDog = append(sliceDog, Dog{
		nickname:       "Кайман",
		age:            2,
		weightKg:       3,
		checkInDate:    "25.11.2022",
		stayShelerName: sliceShelter[1],
	})
	varChooseNameDog := 0
	varPickUpDog := 0

	for {
		clearConsole()

		fmt.Println("Собаки, которых вы могли бы забрать из приюта\nЧтобы узнать больше о питомце введите его номер")

		for index, value := range sliceDog {
			fmt.Println(index+1, value.nickname)

		}

		println("Введите номер собаки, о которой хотите узнать больше:\n")
		fmt.Scan(&varChooseNameDog)
		varChooseNameDog = varChooseNameDog - 1
		clearConsole()

		sliceDog[varChooseNameDog].printDogInfo()

		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("Забрать из приюта?\n1 Да, забрать из приюта\n2 Нет, смотреть других собак")
		fmt.Scan(&varPickUpDog)

		if varPickUpDog == 1 {
			clearConsole()

			sliceShelter[varChooseNameDog].printShelterInfo()

			break

		}
	}

}
