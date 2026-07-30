package main

type Dog struct {
	nickname    string
	age         int
	weightKg    float64
	checkInDate string
	shelter     *Shelter
	policlinic  *Policlinic
}

type Shelter struct {
	numberShelter int
	address       string
	number        string
	workingTime   string
	listPets      map[string]*Dog
}

type Policlinic struct {
	numberClinic int
	address      string
	phoneNumber  string
	workingTime  string
	listPatients map[string]*Dog
}
