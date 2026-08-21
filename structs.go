package main

type Dog struct {
	nickname    string
	age         int
	weightKg    float64
	checkInDate string
}

type Shelter struct {
	numberShelter string
	address       string
	number        string
	workingTime   string
	listPets      map[string]*Dog
}

func (s *Shelter) AddDogs(dog ...Dog) {
	for _, d := range dog {
		s.listPets[d.nickname] = &d
	}
}

func (s *Shelter) RemoveDogs(dog ...Dog) {
}

type Policlinic struct {
	numberClinic string
	address      string
	phoneNumber  string
	workingTime  string
	listPatients map[string]*Dog
}

func (p *Policlinic) AddDogs(dog ...Dog) {
	for _, d := range dog {
		p.listPatients[d.nickname] = &d
	}
}

func (p *Policlinic) RemoveDogs(dog ...Dog) {
}
