package main

import "fmt"

//func clearConsole() {
//	fmt.Print("\033[H\033[2J")
//	println()
//	println()
//}

type Dog struct {
	nickname    string
	age         int
	weightKg    float64
	checkInDate string
	shelter     *Shelter
	policlinic  *Policlinic
}

//func (d Dog) printDogInfo() {
//
//	fmt.Println("Кличка:", d.nickname)
//	fmt.Println("Возраст, лет:", d.age)
//	fmt.Println("Вес, кг:", d.weightKg)
//	fmt.Println("Когда попал в приют:", d.checkInDate)
//	fmt.Println("К какому шелтеру относится:", d.shelter.name)
//
//}

type Shelter struct {
	numberShelter int
	address       string
	number        string
	workingTime   string
	listPets      []*Dog
}

//func (s Shelter) printShelterInfo() {
//	fmt.Println(s.name)
//	fmt.Println(s.address)
//	fmt.Println(s.phoneNumber)
//	fmt.Println(s.workingTime)
//}

type Policlinic struct {
	numberClinic int
	address      string
	phoneNumber  string
	workingTime  string
	listPatients []*Dog
}

//func inputChooseNameDog(lenSliceDog int) (intChooseNameDog int) {
//	for {
//		_, err := fmt.Scan(&intChooseNameDog)
//		if err != nil {
//			fmt.Println("Ошибка ввода")
//			continue
//		}
//		if intChooseNameDog > 0 && intChooseNameDog <= lenSliceDog {
//			break
//		} else {
//			fmt.Println("Ошибка диапазона")
//		}
//
//	}
//	return intChooseNameDog
//}

//func inputPickUpDog() (intPickUpDog int) {
//	for {
//		_, err := fmt.Scan(&intPickUpDog)
//		if err != nil {
//			fmt.Println("Ошибка ввода")
//			continue
//		}
//		if intPickUpDog != 1 && intPickUpDog != 2 {
//			fmt.Println("Ошибка диапазона")
//		} else {
//			break
//		}
//	}
//	return intPickUpDog
//}

func main() {

	policlinics := CreatePoliclinics()
	shelters := CreateShelters()
	dogs := CreateDogs(shelters, policlinics)

	fmt.Println(policlinics[0])
	fmt.Println(shelters)
	fmt.Println(dogs)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(*policlinics[1].listPatients[0])

	//for {
	//	clearConsole()
	//	fmt.Println("Собаки, которых вы могли бы забрать из приюта\nЧтобы узнать больше о питомце введите его номер")
	//	for index, value := range dogs {
	//		fmt.Println(index+1, value.nickname)
	//
	//	}
	//
	//	println("Введите номер собаки, о которой хотите узнать больше:\n")
	//	intChooseNameDog = inputChooseNameDog((len(dogs))) - 1
	//	clearConsole()
	//
	//	dogs[intChooseNameDog].printDogInfo()
	//	fmt.Println()
	//	fmt.Println()
	//	fmt.Println()
	//	fmt.Println("Забрать из приюта?\n1 Да, забрать из приюта\n2 Нет, смотреть других собак")
	//	intPickUpDog = inputPickUpDog()
	//	if intPickUpDog == 1 {
	//		clearConsole()
	//		dogs[intChooseNameDog].shelter.printShelterInfo()
	//		break
	//	}
	//
	//}

}
