package main

type Plants struct {
	Name  string
	Count int
}

func AddPlantsInformation() []Plants {
	return []Plants{
		{Name: "Фикус", Count: 3},
		{Name: "Кактус", Count: 10},
		{Name: "Орхидея", Count: 1},
		{Name: "Монстера", Count: 2},
		{Name: "Ромашка", Count: 5},
	}
}
