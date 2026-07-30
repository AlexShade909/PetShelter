package main

func CreatePoliclinics() []Policlinic {
	policlinics := []Policlinic{
		{
			name:         "Zoo-clinic-1",
			address:      "Mira 1",
			number:       "+37529-566-13-54",
			workingTime:  "10:00-23:00",
			listPatients: []*Dog{},
		}, {
			name:         "Zoo-Pomoshch",
			address:      "Lenina 133",
			number:       "+37529-644-55-71",
			workingTime:  "09:00-24:00",
			listPatients: []*Dog{},
		},
	}
	return policlinics
}

func CreateShelters() []Shelter {
	shelters := []Shelter{
		{
			name:        "Приют 'Шанс'",
			address:     "Пятруся Глебки 17",
			number:      "+375 29 511-22-13",
			workingTime: "10:00 - 22:00",
			listPets:    []*Dog{},
		}, {
			name:        "Приют 'Привет'",
			address:     "Мстислава Чудотворца 4/1",
			number:      "+375 12 544-65-45",
			workingTime: "11:00 - 21:15",
			listPets:    []*Dog{},
		},
	}
	return shelters
}

func CreateDogs(shelters []Shelter, policlinics []Policlinic) []Dog {
	dogs := []Dog{
		{
			nickname:    "Чарли",
			age:         3,
			weightKg:    7.1,
			checkInDate: "05.02.2025",
			shelter:     &shelters[0],
			policlinic:  &policlinics[0],
		}, {
			nickname:    "Вил",
			age:         4,
			weightKg:    17,
			checkInDate: "15.03.2025",
			shelter:     &shelters[1],
			policlinic:  &policlinics[1],
		}, {
			nickname:    "Кайман",
			age:         2,
			weightKg:    3,
			checkInDate: "25.11.2022",
			shelter:     &shelters[1],
			policlinic:  &policlinics[1],
		},
	}

	for i := range dogs {
		dogs[i].shelter.listPets = append(dogs[i].shelter.listPets, &dogs[i])
		dogs[i].policlinic.listPatients = append(dogs[i].policlinic.listPatients, &dogs[i])
	}

	return dogs
}
