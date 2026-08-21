package main

import "fmt"

type PoliclinicsStorage struct {
	table map[string]Policlinic
}

func (p *PoliclinicsStorage) Find(key string) Policlinic {
	pc, ok := p.table[key]
	if !ok {
		fmt.Println("Err! not found policlinic")
	}
	return pc
}

func (p *PoliclinicsStorage) Add(policlinics ...Policlinic) {
	if p.table == nil {
		p.table = make(map[string]Policlinic)
	}
	for _, pc := range policlinics {
		p.table[pc.numberClinic] = pc
	}
}

func (p *PoliclinicsStorage) Remove(key string) {
	delete(p.table, key)
}

//type SheltersStorage struct {
//	table map[string]Shelter
//}
//
//func (s *SheltersStorage) Find(key string) Shelter {
//}
//
//func (s *SheltersStorage) Add(shelter ...Shelter) {
//}
//
//func (s *SheltersStorage) Remove(shelter Shelter) {
//}

type Facade struct {
	PoliclinicsStorage PoliclinicsStorage
	//SheltersStorage    SheltersStorage
}

func main() {
	policlinicsStorage := PoliclinicsStorage{}
	policlinicsStorage.Add(CreatePoliclinics()...)

	fmt.Println(policlinicsStorage.Find("1"))
	policlinicsStorage.Remove("123")
	fmt.Println(policlinicsStorage.Find("1"))

	/*
			sheltersStorage := SheltersStorage{}
			sheltersStorage.Add(CreateShelters()...)

		fmt.Println(sheltersStorage.)
				dogs := CreateDogs()

				policlinnic1 := policlinicsStorage.Find("1")
				policlinnic2 := policlinicsStorage.Find("2")
				shelter1 := sheltersStorage.Find("1")
				shelter2 := sheltersStorage.Find("2")

				policlinnic1.AddDogs(dogs[0], dogs[1])
				policlinnic2.AddDogs(dogs[2])

				shelter1.AddDogs(dogs[0], dogs[2])
				shelter2.AddDogs(dogs[1])


					facade := Facade{
						PoliclinicsStorage: policlinicsStorage,
						SheltersStorage:    sheltersStorage,
					}

					flag := true

					for flag {
						choice := readMenuChoice("1. Выбрать собаку\n2. Добавить собаку\n3. Выход\n ", 1, 3)
						switch choice {
						case 1:
							fmt.Println("Выбрать собаку")
							flag = scenarioTakeDog(facade)
						case 2:
							fmt.Println("Добавить собаку")
							flag = scenarioAddDog(facade)
						case 3:
							fmt.Println("Выход")
							return
						default:
							fmt.Println("неизвестная команда")
							//такой вариант исключён логикой функции readMenuChoice
						}
					}

	*/
}
