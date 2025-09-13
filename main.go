package main

import "fmt"

func main() {
	numAnts, Rooms, Links, StartRoom, EndRoom, err := ParseFile("examples/file3.txt")
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
	fmt.Println("Paths found:")
	for i, p := range paths {
		fmt.Printf("Path %d: %v\n", i+1, p.Rooms)
	}

	antsOnPath := distributeAnts(paths, numAnts)
	fmt.Println("\nAnt distribution:")
	for i, n := range antsOnPath {
		fmt.Printf("Path %d: %d ants\n", i+1, n)
	}

	antNum := 1
	fmt.Println("\nAnt assignments:")
	for i, n := range antsOnPath {
		for j := 0; j < n; j++ {
			fmt.Printf("Ant %d -> Path %d (%v)\n", antNum, i+1, paths[i].Rooms)
			antNum++
		}
	}
}
