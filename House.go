package main

type House struct {
	Length               float64
	Width                float64
	Height               float64
	FamilyInformation    []Family
	PetsInformation      []Pets
	RoomsInformation     []Rooms
	PlantsInformation    []Plants
	TablewareInformation []Tableware
}

func createHouse() House {
	return House{
		Length:               6.8,
		Width:                5,
		Height:               2.9,
		FamilyInformation:    AddFamilyInformation(),
		PetsInformation:      AddPetsInformation(),
		RoomsInformation:     AddRoomsInformation(),
		PlantsInformation:    AddPlantsInformation(),
		TablewareInformation: AddTablewareInformation(),
	}
}
