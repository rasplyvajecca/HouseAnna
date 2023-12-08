package house

import (
	"HouseAnna/House/family"
	"HouseAnna/House/pets"
	"HouseAnna/House/plants"
	"HouseAnna/House/rooms"
	"HouseAnna/House/tablewear"
	"fmt"
)

type House struct {
	Length               float64
	Width                float64
	Height               float64
	FamilyInformation    []family.Family
	PetsInformation      []pets.Pets
	RoomsInformation     []rooms.Rooms
	PlantsInformation    []plants.Plants
	TablewareInformation []tablewear.Tableware
}

func CreateHouse() House {
	return House{
		Length:               6.8,
		Width:                5,
		Height:               2.9,
		FamilyInformation:    family.AddFamilyInformation(),
		PetsInformation:      pets.AddPetsInformation(),
		RoomsInformation:     rooms.AddRoomsInformation(),
		PlantsInformation:    plants.AddPlantsInformation(),
		TablewareInformation: tablewear.AddTablewareInformation(),
	}
}

func DescribeHouse(house House) {
	fmt.Printf("Описание дома:\n")
	fmt.Printf("Длина: %.2f м\n", house.Length)
	fmt.Printf("Ширина: %.2f м\n", house.Width)
	fmt.Printf("Высота:%.2f м\n", house.Height)

	fmt.Println("Описание семьи:")
	for _, member := range house.FamilyInformation {
		fmt.Printf("- %s: %s\n", member.Member, member.Name)
	}

	fmt.Println("Описание животных:")
	for _, pet := range house.PetsInformation {
		fmt.Printf("- %s: %s\n", pet.Member, pet.Name)
	}

	fmt.Println("Описание комнат:")
	for _, room := range house.RoomsInformation {
		fmt.Printf("- %s: %.2f кв. м\n", room.Name, room.Size)
	}

	fmt.Println("Описание растений:")
	for _, plant := range house.PlantsInformation {
		fmt.Printf("- %s: %d шт.\n", plant.Name, plant.Count)
	}

	fmt.Println("Описание столовых приборов:")
	for _, tableware := range house.TablewareInformation {
		fmt.Printf("- %s: %s, %d шт.\n", tableware.Name, tableware.Material, tableware.Count)
	}
}
