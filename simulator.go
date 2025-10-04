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

func SimulateAnts(finalPaths [][]string, antsOnPath []int) {
	type Ant struct {
		ID   int
		Path []string
		Pos  int
	}

	// Initialize all ants
	ants := []Ant{}
	antID := 1
	for i, n := range antsOnPath {
		for j := 0; j < n; j++ {
			ants = append(ants, Ant{
				ID:   antID,
				Path: finalPaths[i],
				Pos:  0,
			})
			antID++
		}
	}

	// Turn by turn simulation
	for {
		allAtEnd := true
		movesThisTurn := []string{}

		for i := range ants {
			if ants[i].Pos < len(ants[i].Path)-1 {
				ants[i].Pos++
				movesThisTurn = append(movesThisTurn,
					fmt.Sprintf("L%d-%s", ants[i].ID, ants[i].Path[ants[i].Pos]))
				allAtEnd = false
			}
		}

		if len(movesThisTurn) > 0 {
			fmt.Println(strings.Join(movesThisTurn, " "))
		}

		if allAtEnd {
			break
		}
	}
}
