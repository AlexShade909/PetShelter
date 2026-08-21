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
	numberShelter string
	address       string
	number        string
	workingTime   string
}

type Policlinic struct {
	numberClinic string
	address      string
	phoneNumber  string
	workingTime  string
}
