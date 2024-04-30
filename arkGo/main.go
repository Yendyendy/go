package main

import (
	"fmt"
	"math/rand"
)

type Animal interface {
	move() string
	// eat() string
}

type Cow struct{}
type Duck struct{}
type Snake struct{}
type Bird struct{}

func (c Cow) String() string {
	return "Cow$kj"
}
func (c Cow) move() string {
	return c.String() + " move"
}

func (d Duck) String() string {
	return "Duck$"
}
func (d Duck) move() string {
	return d.String() + " move"
}

func (s Snake) String() string {
	return "Snake$"
}
func (s Snake) move() string {
	return s.String() + " move"
}

func (b Bird) String() string {
	return "Bird$$"
}
func (b Bird) move() string {
	return b.String() + " move"
}
func main() {
	fmt.Println("Hello worlds")

	ark := []Animal{Duck{}, Cow{}, Bird{}, Snake{}}

	for i := 1; i < 4; i++ {
		for j := 0; j < 24; j++ {
			n := rand.Intn(4)
			animal := ark[n]
			fmt.Printf("Dia: %d- hora: %d -> %s.\n", i, j, animal.move())
		}
	}

}
