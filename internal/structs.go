package internal

type Dog struct {
	Nickname    string
	Age         int
	WeightKg    float64
	CheckInDate string
	Shelter     *Shelter
	Policlinic  *Policlinic
}

type Shelter struct {
	NumberShelter string
	Address       string
	Number        string
	WorkingTime   string
}

type Policlinic struct {
	NumberClinic string
	Address      string
	PhoneNumber  string
	WorkingTime  string
}
