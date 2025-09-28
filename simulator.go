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

	// 2. Turn loop
	for !allAntsAtEnd(ants) {
		movedThisTurn := []string{} // keep track for printing

		for i := range ants {
			nextPos := ants[i].Pos + 1
			if nextPos < len(ants[i].Path) {
				nextRoom := ants[i].Path[nextPos]

				// check if next room is free (except end)
				if isRoomFree(nextRoom, ants) {
					ants[i].Pos = nextPos
					movedThisTurn = append(movedThisTurn, fmt.Sprintf("L%d-%s", ants[i].ID, nextRoom))
				}
			}
		}

		// print the turn
		if len(movedThisTurn) > 0 {
			fmt.Println(strings.Join(movedThisTurn, " "))
		}
	}
}

func SimulateAntMovement(paths [][]string, antDistribution [][]int) string {
	var finalResult string
	type AntPosition struct {
		ant  int
		path int
		step int
	}

	var antPositions []AntPosition
	for pathIndex, ants := range antDistribution {
		for _, ant := range ants {
			antPositions = append(antPositions, AntPosition{ant, pathIndex, 0})
		}
	}
	for len(antPositions) > 0 {
		var moves []string
		var newPositions []AntPosition
		usedLinks := make(map[string]bool)

		for _, pos := range antPositions {
			if pos.step < len(paths[pos.path])-1 {
				currentRoom := paths[pos.path][pos.step]
				nextRoom := paths[pos.path][pos.step+1]
				link := currentRoom + "-" + nextRoom
				if !usedLinks[link] {
					moves = append(moves, fmt.Sprintf("L%d-%s", pos.ant, nextRoom))
					newPositions = append(newPositions, AntPosition{pos.ant, pos.path, pos.step + 1})
					usedLinks[link] = true
				} else {
					newPositions = append(newPositions, pos)
				}
			}
		}
		if len(moves) > 0 {
			finalResult += strings.Join(moves, " ")
			finalResult += "\n"
		}
		antPositions = newPositions
	}
	return finalResult
}
