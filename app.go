package main

import (
	"houseAnna/house/home"
)

func main() {
	houseAnna := home.CreateHome()
	home.DescribeHome(houseAnna)
}
