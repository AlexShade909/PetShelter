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
	stayShelerName Shelter
}

func (d Dog) printDogInfo() {

	fmt.Println("Кличка:", d.nickname)
	fmt.Println("Номер в базе:", d.numberInBase)
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
		numberInBase:   1,
		nickname:       "Чарли",
		age:            3,
		weightKg:       7,
		checkInDate:    "05.02.2025",
		stayShelerName: sliceShelter[0],
	})

	sliceDog = append(sliceDog, Dog{
		numberInBase:   2,
		nickname:       "Вил",
		age:            4,
		weightKg:       17,
		checkInDate:    "15.03.2025",
		stayShelerName: sliceShelter[1],
	})
	sliceDog = append(sliceDog, Dog{
		numberInBase:   3,
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
		fmt.Println(sliceDog[0].numberInBase, sliceDog[0].nickname)
		fmt.Println(sliceDog[1].numberInBase, sliceDog[1].nickname)
		fmt.Println(sliceDog[2].numberInBase, sliceDog[2].nickname)

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
