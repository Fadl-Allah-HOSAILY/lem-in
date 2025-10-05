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

func SimulateAnts(paths [][]string, antsOnPath []int) {
	type Ant struct {
		ID   int
		Path []string
		Pos  int
	}

	var ants []Ant
	antID := 1
	totalAnts := 0
	for _, c := range antsOnPath {
		totalAnts += c
	}

	// Holds ants currently moving
	activeAnts := []Ant{}

	for turn := 1; len(ants) < totalAnts || len(activeAnts) > 0; turn++ {
		var moves []string

		// Move existing ants
		newActive := []Ant{}
		for _, ant := range activeAnts {
			if ant.Pos < len(ant.Path)-1 {
				ant.Pos++
				moves = append(moves, fmt.Sprintf("L%d-%s", ant.ID, ant.Path[ant.Pos]))
				if ant.Pos < len(ant.Path)-1 {
					newActive = append(newActive, ant)
				}
			}
		}
		activeAnts = newActive

		// Launch one new ant per path (if any left)
		for i, count := range antsOnPath {
			if count > 0 && antID <= totalAnts {
				path := paths[i]
				ant := Ant{ID: antID, Path: path, Pos: 0}
				ant.Pos++
				moves = append(moves, fmt.Sprintf("L%d-%s", ant.ID, path[ant.Pos]))
				activeAnts = append(activeAnts, ant)
				ants = append(ants, ant)
				antsOnPath[i]--
				antID++
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}
