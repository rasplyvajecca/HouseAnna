package main

type Tableware struct {
	Name     string
	Material string
	Count    int
}

func AddTablewareInformation() []Tableware {
	return []Tableware{
		{Name: "Ложка", Material: "Серебро", Count: 5},
		{Name: "Вилка", Material: "Серебро", Count: 5},
		{Name: "Нож", Material: "Золото", Count: 1},
		{Name: "Венчик", Material: "Платик", Count: 1},
		{Name: "Чайная ложка", Material: "Серебро", Count: 5},
	}
}
