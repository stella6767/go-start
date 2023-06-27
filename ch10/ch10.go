package main

import (
	ch9 "awesomeProject/ch9"
	"fmt"
)

/**
golang에서 oop란

*/

type Bicycle struct {
}

func (b Bicycle) Run() {
	fmt.Println("Bicycle Run")
}

type MountainBike struct {
	Bicycle
}

type RoadBike struct {
	Bicycle
}

// 오버라이딩을 이런식으록 구현하나? 번거로운데?

func (m MountainBike) Run() {
	fmt.Println("Mountain Bicycle Run")
}

func (m RoadBike) Run() {
	fmt.Println("Road Bicycle Run")
}

func main() {

	person := new(ch9.Person)
	fmt.Println(person)

	rBike := RoadBike{}
	rBike.Run()
	mBike := MountainBike{}
	mBike.Run()
}
