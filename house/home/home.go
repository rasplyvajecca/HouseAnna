package home

import (
	"fmt"
	"houseAnna/house/family"
	"houseAnna/house/pets"
	"houseAnna/house/plants"
	"houseAnna/house/rooms"
	"houseAnna/house/tableware"
)

type Home struct {
	LengthMeters         float64
	WidthMeters          float64
	HeightMeters         float64
	FamilyInformation    []family.Family
	PetsInformation      []pets.Pets
	RoomsInformation     []rooms.Rooms
	PlantsInformation    []plants.Plants
	TablewareInformation []tableware.Tableware
}

func CreateHome() Home {
	return Home{
		LengthMeters:         6.8,
		WidthMeters:          5,
		HeightMeters:         2.9,
		FamilyInformation:    family.AddFamilyInformation(),
		PetsInformation:      pets.AddPetsInformation(),
		RoomsInformation:     rooms.AddRoomsInformation(),
		PlantsInformation:    plants.AddPlantsInformation(),
		TablewareInformation: tableware.AddTablewareInformation(),
	}
}

func DescribeHome(home Home) {
	fmt.Printf("Описание дома:\n")
	printHomeDimensions(home)
	printFamilyInformation(home.FamilyInformation)
	printPetsInformation(home.PetsInformation)
	printRoomsInformation(home.RoomsInformation)
	printPlantsInformation(home.PlantsInformation)
	printTablewareInformation(home.TablewareInformation)
}

func printHomeDimensions(home Home) {
	fmt.Printf("Длина: %.2f м\n", home.LengthMeters)
	fmt.Printf("Ширина: %.2f м\n", home.WidthMeters)
	fmt.Printf("Высота: %.2f м\n", home.HeightMeters)
}

func printFamilyInformation(familyInfo []family.Family) {
	fmt.Println("Описание семьи:")
	for _, member := range familyInfo {
		fmt.Printf("- %s: %s\n", member.Member, member.Name)
	}
}

func printPetsInformation(petsInfo []pets.Pets) {
	fmt.Println("Описание животных:")
	for _, pet := range petsInfo {
		fmt.Printf("- %s: %s\n", pet.Member, pet.Name)
	}
}

func printRoomsInformation(roomsInfo []rooms.Rooms) {
	fmt.Println("Описание комнат:")
	for _, room := range roomsInfo {
		fmt.Printf("- %s: %.2f кв. м\n", room.Name, room.Size)
	}
}

func printPlantsInformation(plantsInfo []plants.Plants) {
	fmt.Println("Описание растений:")
	for _, plant := range plantsInfo {
		fmt.Printf("- %s: %d шт.\n", plant.Name, plant.Count)
	}
}

func printTablewareInformation(tablewareInfo []tableware.Tableware) {
	fmt.Println("Описание столовых приборов:")
	for _, object := range tablewareInfo {
		fmt.Printf("- %s: %s, %d шт.\n", object.Name, object.Material, object.Count)
	}
}
