package main

import "fmt"

func main() {
	numAnts, Rooms, Links, StartRoom, EndRoom, _, err := ParseFile("examples/file0.txt")
	if err != nil {
		panic(err)
	}

	graph := &Graph{
		Rooms: Rooms,
		Links: Links,
		Start: StartRoom,
		End:   EndRoom,
	}

	paths := FindMultiplePaths(graph)
	fmt.Println(paths)
	antsOnPath := distributeAnts(paths, numAnts)
	SimulateAntMovement(paths, antsOnPath)
}
