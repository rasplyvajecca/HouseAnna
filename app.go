package main

import (
	"HouseAnna/House/house"
)

func main() {
	houseAnna := house.CreateHouse()
	house.DescribeHouse(houseAnna)
}
