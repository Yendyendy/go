package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type SpaceLineName int

const (
	VirginGalactic  SpaceLineName = 0
	Spacex          SpaceLineName = 1
	SpaceAdventures SpaceLineName = 2
)

type TripType int

const (
	RoundTrip TripType = 0
	OneWay    TripType = 1
)

type SpaceLine struct {
	name     SpaceLineName
	days     int
	tripType TripType
	price    int
}

const width = 20

func main() {
	fmt.Println("Wellcome to SpaceGo!\n")

	createTableHead()
	for i := 0; i < 10; i++ {
		printRow(newSpaceLine())

	}
}

func createTableHead() {

	fmt.Printf("%-*s %-*s %-*s %-*s\n", width, "Space Line", width, "Days", width, "Trip type", width, "Price")

	h := strings.Repeat("=", width*4)

	fmt.Println(h)
}

func printRow(spaceLine SpaceLine) {
	fmt.Printf("%-*s %-*d %-*s %-*d\n",
		width, spaceLineToString(spaceLine.name),
		width, spaceLine.days,
		width, tripTypeToString(spaceLine.tripType),
		width, spaceLine.price)
}

func newSpaceLine() SpaceLine {
	//get space line
	spaceLineName := SpaceLineName(randRange(int(VirginGalactic), int(SpaceAdventures)))
	days := getDays()
	tripType := TripType(randRange(int(RoundTrip), int(OneWay)))
	price := randRange(36, 50)

	if tripType == 0 {
		price *= 2
	}

	return SpaceLine{
		spaceLineName,
		days,
		tripType,
		price,
	}
}

func randRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func getDays() int {
	const distance = 62100000
	speed := randRange(16, 30)

	seconds := distance / speed
	return seconds / (60 * 60 * 24) //days

}

func spaceLineToString(id SpaceLineName) string {
	switch id {
	case VirginGalactic:
		return "Virgin Galactic"
	case Spacex:
		return "SpaceX"
	case SpaceAdventures:
		return "Space Adventures"
	default:
		return "error"
	}
}

func tripTypeToString(id TripType) string {
	switch id {
	case RoundTrip:
		return "Round-trip"
	case OneWay:
		return "One-way"
	default:
		return "error"
	}
}
