package main

import (
	"fmt"
	"strings"
)

type Ant struct {
	ID   int
	Path []string
	Pos  int // index in Path
}

func allAntsAtEnd(ants []Ant) bool {
	for _, ant := range ants {
		if ant.Pos != len(ant.Path)-1 { // still moving
			return false
		}
	}
	return true
}

func moveAnt(ant *Ant) (moved bool, room string) {
	if ant.Pos < len(ant.Path)-1 {
		ant.Pos++
		return true, ant.Path[ant.Pos]
	}
	return false, ""
}

func moveAnts(ants []Ant) []string {
	moves := []string{}
	for i := range ants {
		if moved, room := moveAnt(&ants[i]); moved {
			moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].ID, room))
		}
	}
	return moves
}

func printMoves(moves []string) {
	if len(moves) > 0 {
		fmt.Println(strings.Join(moves, " "))
	}
}

// ---------------------- Main Simulation ----------------------

func SimulateAnts(paths []Path, antsOnPath []int) {
	// 1. Initialize ants
	ants := []Ant{}
	antID := 1
	for i, n := range antsOnPath {
		for j := 0; j < n; j++ {
			ants = append(ants, Ant{
				ID:   antID,
				Path: paths[i].Rooms,
				Pos:  0,
			})
			antID++
		}
	}

	// 2. Turn loop
	for !allAntsAtEnd(ants) {
		moves := moveAnts(ants)
		printMoves(moves)
	}
}
