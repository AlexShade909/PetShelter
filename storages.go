package main

import "fmt"

type PoliclinicsStorage struct {
	table map[string]Policlinic
}

func (p *PoliclinicsStorage) Find(key string) Policlinic {
	pc, ok := p.table[key]
	if !ok {
		fmt.Println("Err! not found policlinic")
	}
	return pc
}

func (p *PoliclinicsStorage) Add(policlinics ...Policlinic) {
	if p.table == nil {
		p.table = make(map[string]Policlinic)
	}
	for _, pc := range policlinics {
		p.table[pc.numberClinic] = pc
	}
}

func (p *PoliclinicsStorage) Remove(key string) {
	delete(p.table, key)
}

type SheltersStorage struct {
	table map[string]Shelter
}

func (s *SheltersStorage) Find(key string) Shelter {
	shstr, ok := s.table[key]
	if !ok {
		fmt.Println("Err! Not found shelter")
	}
	return shstr
}
func (s *SheltersStorage) Add(shelter ...Shelter) {
	if s.table == nil {
		s.table = make(map[string]Shelter)
	}
	for _, shstr := range shelter {
		s.table[shstr.numberShelter] = shstr
	}
}
func (s *SheltersStorage) Remove(key string) {
	delete(s.table, key)
}

type FacadeStorage struct {
	PoliclinicsStorage PoliclinicsStorage
	//SheltersStorage    SheltersStorage
}
