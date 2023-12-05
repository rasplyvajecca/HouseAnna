package main

type Pets struct {
	Member string
	Name   string
}

func AddPetsInformation() []Pets {
	return []Pets{
		{Member: "Кот", Name: "Мурзик"},
		{Member: "Кошка", Name: "Мурка"},
		{Member: "Кошка", Name: "Зоя"},
		{Member: "Кролик", Name: "Джо"},
		{Member: "Хомяк", Name: "Хома"},
	}
}
