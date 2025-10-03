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

func isRoomFree(room string, ants []Ant) bool {
	if room == ants[0].Path[len(ants[0].Path)-1] { // End room
		return true
	}
	for _, a := range ants {
		if a.Path[a.Pos] == room {
			return false
		}
	}
	return true
}

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

	// 2. Turn by turn
	for !allAntsAtEnd(ants) {
		movesThisTurn := []string{}

		for i := range ants {
			if ants[i].Pos < len(ants[i].Path)-1 {
				ants[i].Pos++
				movesThisTurn = append(movesThisTurn,
					fmt.Sprintf("L%d-%s", ants[i].ID, ants[i].Path[ants[i].Pos]))
			}
		}

		if len(movesThisTurn) > 0 {
			fmt.Println(strings.Join(movesThisTurn, " "))
		}
	}
}
