package main

import "fmt"

func main() {

	policlinics := CreatePoliclinics()
	shelters := CreateShelters()
	dogs := CreateDogs(shelters, policlinics)

	fmt.Println(dogs)
	scenarioTakeDog()
}
