package main

func main() {
	numAnts, Rooms, Links, StartRoom, EndRoom, _, err := ParseFile("examples/big_7.txt")
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
	antsOnPath := distributeAnts(paths, numAnts)
	SimulateAnts(paths, antsOnPath)
}
