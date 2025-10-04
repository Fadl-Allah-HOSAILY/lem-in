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

// simulate ants turn by turn and print moves
func allAntsAtEnd(ants []Ant) bool {
	for _, ant := range ants {
		if ant.Pos < len(ant.Path)-1 {
			return false
		}
	}
	return true
}

func SimulateAnts(paths [][]string, antsOnPath []int) {
	// 1. Initialize ants
	ants := []Ant{}
	antID := 1
	for i, n := range antsOnPath {
		for j := 0; j < n; j++ {
			ants = append(ants, Ant{
				ID:   antID,
				Path: paths[i],
				Pos:  0, // starting at the start room
			})
			antID++
		}
	}

	// 2. Turn by turn
	for !allAntsAtEnd(ants) {
		movesThisTurn := []string{}
		occupied := map[string]bool{} // intermediate rooms occupied this turn

		for i := range ants {
			if ants[i].Pos < len(ants[i].Path)-1 {
				nextRoom := ants[i].Path[ants[i].Pos+1]

				// move if room is free or it's the end
				if !occupied[nextRoom] || nextRoom == ants[i].Path[len(ants[i].Path)-1] {
					ants[i].Pos++
					movesThisTurn = append(movesThisTurn,
						fmt.Sprintf("L%d-%s", ants[i].ID, nextRoom))
					if nextRoom != ants[i].Path[len(ants[i].Path)-1] {
						occupied[nextRoom] = true
					}
				}
			}
		}

		if len(movesThisTurn) > 0 {
			fmt.Println(strings.Join(movesThisTurn, " "))
		}
	}
}
