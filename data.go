package main

func CreatePoliclinics() []Policlinic {
	policlinics := []Policlinic{
		{
			numberClinic: 1,
			address:      "Мира 1",
			phoneNumber:  "+37529-566-13-54",
			workingTime:  "10:00-23:00",
			listPatients: make(map[string]*Dog),
		}, {
			numberClinic: 12,
			address:      "Ленина 133",
			phoneNumber:  "+37529-644-55-71",
			workingTime:  "09:00-24:00",
			listPatients: make(map[string]*Dog),
		},
	}
	return policlinics
}

func CreateShelters() []Shelter {
	shelters := []Shelter{
		{
			numberShelter: 1,
			address:       "Пятруся Глебки 17",
			number:        "+375 29 511-22-13",
			workingTime:   "10:00 - 22:00",
			listPets:      make(map[string]*Dog),
		}, {
			numberShelter: 2,
			address:       "Мстислава Чудотворца 4/1",
			number:        "+375 12 544-65-45",
			workingTime:   "11:00 - 21:15",
			listPets:      make(map[string]*Dog),
		},
	}
	return shelters
}

func AddDog(dogs map[string]*Dog, nickname string, age int, weightKg float64, checkInDate string, shelter *Shelter, policlinic *Policlinic) *Dog {
	d := &Dog{
		nickname:    nickname,
		age:         age,
		weightKg:    weightKg,
		checkInDate: checkInDate,
		shelter:     shelter,
		policlinic:  policlinic,
	}

	dogs[nickname] = d
	shelter.listPets[nickname] = d
	policlinic.listPatients[nickname] = d

	return d
}

func RemoveDog(dogs map[string]*Dog, nickname string) bool {
	d, ok := dogs[nickname]
	if !ok {
		return false
	}

	delete(d.shelter.listPets, nickname)
	delete(d.policlinic.listPatients, nickname)
	delete(dogs, nickname)

	return true
}

func CreateDogs(shelters []Shelter, policlinics []Policlinic) map[string]*Dog {
	dogs := make(map[string]*Dog)

	AddDog(dogs, "Чарли", 3, 7.1, "05.02.2025", &shelters[0], &policlinics[0])
	AddDog(dogs, "Вил", 4, 17, "15.03.2025", &shelters[1], &policlinics[1])
	AddDog(dogs, "Кайман", 2, 3, "25.11.2022", &shelters[1], &policlinics[1])

	return dogs
}
