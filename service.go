package main

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

func FindDog(dogs map[string]*Dog, nickname string) (*Dog, bool) {
	d, ok := dogs[nickname]
	return d, ok
}
