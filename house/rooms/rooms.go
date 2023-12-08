package rooms

type Rooms struct {
	Name string
	Size float64
}

func AddRoomsInformation() []Rooms {
	return []Rooms{
		{Name: "Кухня", Size: 5},
		{Name: "Ванная", Size: 3},
		{Name: "Спальня", Size: 10},
		{Name: "Зал", Size: 20},
		{Name: "Детская комната", Size: 50},
	}
}
