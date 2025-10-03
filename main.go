package main

import "fmt"

/*
##start
##end FIX

A 5 1
c 5 1

h 4 6
A 5
c 8 1

dup link e-e
*/

func main() {
	numAnts, Rooms, Links, StartRoom, EndRoom, err := ParseFile("examples/file0.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	graph := &Graph{
		Rooms: Rooms,
		Links: Links,
		Start: StartRoom,
		End:   EndRoom,
	}

	paths := FindMultiplePaths(graph)
	distrobution := distributeAnts(paths, numAnts)
	SimulateAnts(paths, distrobution)
}
