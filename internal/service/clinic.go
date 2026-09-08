package service

import "PetShelter/internal/models"

// TODO: Написать сервис поликлинники

type Clinic struct{}

func NewClinic() *Clinic {
	return &Clinic{}
}

func (c *Clinic) GetAllClinics() []models.Policlinic {
	return []models.Policlinic{
		{
			NumberClinic: "Поликлиника 0",
			Address:      "Мира 1",
			PhoneNumber:  "+37529-566-13-54",
			WorkingTime:  "10:00-23:00",
		},
		{
			NumberClinic: "Поликлиника 1",
			Address:      "Ленина 133",
			PhoneNumber:  "+37529-644-55-71",
			WorkingTime:  "09:00-24:00",
		},
	}
}
