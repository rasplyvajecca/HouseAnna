package main

import (
	"HouseAnna/house/home"
)

func main() {
	houseAnna := home.CreateHome()
	home.DescribeHome(houseAnna)
}
