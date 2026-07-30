package main

//func clearConsole() {
//	fmt.Print("\033[H\033[2J")
//	println()
//	println()
//}

//func (s Shelter) printShelterInfo() {
//	fmt.Println(s.name)
//	fmt.Println(s.address)
//	fmt.Println(s.phoneNumber)
//	fmt.Println(s.workingTime)
//}

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

	printDogInfo(dogs, "Вил")
}
