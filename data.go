package main

func CreatePoliclinics() []Policlinic {
	policlinics := []Policlinic{
		{
			numberClinic: "Поликлиника 0",
			address:      "Мира 1",
			phoneNumber:  "+37529-566-13-54",
			workingTime:  "10:00-23:00",
		},
		{
			numberClinic: "Поликлиника 1",
			address:      "Ленина 133",
			phoneNumber:  "+37529-644-55-71",
			workingTime:  "09:00-24:00",
		},
	}
	return policlinics
}

func CreateShelters() []Shelter {
	shelters := []Shelter{
		{
			numberShelter: "Шелтер 0",
			address:       "Пятруся Глебки 17",
			number:        "+375 29 511-22-13",
			workingTime:   "10:00 - 22:00",
		},
		{
			numberShelter: "Шелтер 1",
			address:       "Мстислава Чудотворца 4/1",
			number:        "+375 12 544-65-45",
			workingTime:   "11:00 - 21:15",
		},
	}
	return shelters
}
func CreateDogs(shelter []Shelter, policlinic []Policlinic) map[string]Dog {
	dogs := map[string]Dog{
		"Чарли": {
			nickname:    "Чарли",
			age:         12,
			weightKg:    135,
			checkInDate: "05.02.2025",
			shelter:     &shelter[0],
			policlinic:  &policlinic[0],
		},
		"Спайси": {
			nickname:    "Спайси",
			age:         13,
			weightKg:    15,
			checkInDate: "15.12.2025",
			shelter:     &shelter[1],
			policlinic:  &policlinic[0],
		},
		"Кайман": {
			nickname:    "Кайман",
			age:         643,
			weightKg:    12,
			checkInDate: "05.07.2025",
			shelter:     &shelter[0],
			policlinic:  &policlinic[1],
		},
	}
	return dogs
}
