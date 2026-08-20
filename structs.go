package main

type Dog struct {
	nickname    string
	age         int
	weightKg    float64
	checkInDate string
}

type Shelter struct {
	numberShelter int
	address       string
	number        string
	workingTime   string
	listPets      map[string]*Dog
}

func (s *Shelter) AddDogs(dog ...Dog) {
}

func (s *Shelter) RemoveDogs(dog ...Dog) {
}

type Policlinic struct {
	numberClinic int
	address      string
	phoneNumber  string
	workingTime  string
	listPatients map[string]*Dog
}

func (p *Policlinic) AddDogs(dog ...Dog) {
}

func (p *Policlinic) RemoveDogs(dog ...Dog) {
}
