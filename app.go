package main

import (
	"HouseAnna/house/house"
)

func main() {
	houseAnna := house.CreateHouse()
	house.DescribeHouse(houseAnna)
}
