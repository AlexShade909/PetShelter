Стуктура Собака
имя
возраст
вес
дата регистрации
шелтер пребывания
поликлиника регистрации

- М. Забрать собаку

Стуктура Шелтер
название 
адресс
номер
время работы
список питомцев

- М. Удалить собаку из базы

Структура Поликлиника
название
адресс
номер
время работы
список пациентов


- М. Удалить пациента


type Dog struct {
nickname    string
age         int
weightKg    float64
checkInDate string
shelter     *Shelter
policlinic  *Policlinic

type Shelter struct {
name        string
address     string
number      string
workingTime string
listPets    []Dog
}

type Policlinic struct {
name         string
address      string
number       string
workingTime  string
listPatients []
}