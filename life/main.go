package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width = 80

	height = 15
)

type Universe [height][width]bool

var NeighborsPositions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func main() {
	fmt.Println("Hello life")

	u := NewUniverse()
	u.Seed()
	u.Show()

	for {
		fmt.Scanf(">")
		w := NewUniverse()
		Step(&u, &w)
		u.Show()
	}
}

func NewUniverse() Universe {

	return Universe{}
}

func (u Universe) Show() {
	for _, row := range u {
		for _, col := range row {
			if col {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println()
	}
}

func (u *Universe) Seed() {
	rand.Seed(time.Now().Unix())
	total := width * height

	i := 0
	aux := int(float32(total) * 0.25)
	for i < aux {
		cel := rand.Intn(total)
		fila := cel / width
		col := cel % width
		if !u[fila][col] {
			(*u)[fila][col] = true
			i++
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	return u[(x+height)%height][(y+width)%width]
}

func (u Universe) Neighbors(x, y int) int {

	count := 0

	for _, pos := range NeighborsPositions {
		if u.Alive(x+pos[0], y+pos[1]) {
			count++
		}
	}

	return count
}

func (u Universe) Next(x, y int) bool {

	nNeighbors := u.Neighbors(x, y)

	if u.Alive(x, y) {
		if nNeighbors < 2 {
			return false
		} else if nNeighbors == 2 || nNeighbors == 3 {
			return true
		} else if nNeighbors > 3 {
			return false
		}
	} else {
		if nNeighbors == 3 {
			return true
		}
	}
	return false
}

func Step(u, w *Universe) {
	for i, row := range u {
		for j, _ := range row {
			w[i][j] = u.Next(i, j)
		}
	}

	(*u), (*w) = (*w), (*u)
}
