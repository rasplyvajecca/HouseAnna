package main

type Family struct {
	Member string
	Name   string
}

func AddFamilyInformation() []Family {
	return []Family{
		{Member: "Мама", Name: "Людмила"},
		{Member: "Папа", Name: "Олег"},
		{Member: "Сын", Name: "Макс"},
		{Member: "Дачь", Name: "Мария"},
		{Member: "Бабушка", Name: "Ксения"},
	}
}
