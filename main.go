package main

import (
	"fmt"
	"os"

	"lemin/farm"
	"lemin/pathfinder"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . file.txt")
		return
	}

	file := os.Args[1]

	NumAnts, Rooms, StartRoom, EndRoom, err := ParseFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	farm := farm.Farm{
		Ants:  NumAnts,
		Rooms: Rooms,
		Start: StartRoom,
		End:   EndRoom,
	}

	allPaths := pathfinder.FindAllShortestPaths(farm)

	bestPaths := pathfinder.SelectBestPaths(farm, allPaths)

	nonOverlapPaths := pathfinder.FindNonOverlappingPaths(farm)

	var finalPaths [][]string
	if len(nonOverlapPaths) > len(bestPaths) {
		finalPaths = nonOverlapPaths
	} else {
		finalPaths = bestPaths
	}

	if len(finalPaths) == 0 {
		fmt.Println("No valid paths found!")
		return
	}

	distrobution := distributeAnts(finalPaths, NumAnts)
	SimulateAnts(finalPaths, distrobution)
}
