package main

func AddDog(dogs map[string]Dog, nickname string, age int, weightKg float64, checkInDate string, shelter *Shelter, policlinic *Policlinic) *Dog {
	d := &Dog{
		nickname:    nickname,
		age:         age,
		weightKg:    weightKg,
		checkInDate: checkInDate,
		shelter:     shelter,
		policlinic:  policlinic,
	}
	dogs[nickname] = *d
	return d
}

func RemoveDog(dogs map[string]Dog, nickname string) bool {
	_, ok := dogs[nickname]
	if !ok {
		return false
	}
	delete(dogs, nickname)
	return true
}

func FindDog(dogs map[string]Dog, nickname string) (Dog, bool) {
	d, ok := dogs[nickname]
	return d, ok
}
