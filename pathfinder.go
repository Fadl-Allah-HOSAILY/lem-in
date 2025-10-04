package main

type Graph struct {
	Rooms map[string]Room
	Links map[string][]string
	Start string
	End   string
}

func findNonOverlappingPaths(f *Farm) [][]string {
	var selectedPaths [][]string
	blockedRooms := make(map[string]bool)

	// Sort neighbors by number of links (try neighbors with fewer connections first)
	neighbors := make([]string, len(f.Rooms[f.Start].Links))
	copy(neighbors, f.Rooms[f.Start].Links)

	// Simple sorting by number of links
	for i := 0; i < len(neighbors)-1; i++ {
		for j := i + 1; j < len(neighbors); j++ {
			if len(f.Rooms[neighbors[j]].Links) < len(f.Rooms[neighbors[i]].Links) {
				neighbors[i], neighbors[j] = neighbors[j], neighbors[i]
			}
		}
	}

	// Find path for each neighbor
	for _, neighbor := range neighbors {
		path := bfsShortestPath(f, neighbor, blockedRooms)
		if path != nil {
			selectedPaths = append(selectedPaths, path)

			// Block intermediate rooms from this path (except start and end)
			for i := 1; i < len(path)-1; i++ {
				blockedRooms[path[i]] = true
			}
		}
	}

	return selectedPaths
}

// ----- Optimized BFS to find shortest path avoiding blocked rooms -----
func bfsShortestPath(f *Farm, startNeighbor string, blockedRooms map[string]bool) []string {
	queue := [][]string{{f.Start, startNeighbor}}
	visited := make(map[string]bool)
	visited[f.Start] = true
	visited[startNeighbor] = true

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		last := path[len(path)-1]

		if last == f.End {
			return path
		}

		for _, next := range f.Rooms[last].Links {
			if !visited[next] && !blockedRooms[next] {
				visited[next] = true
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, next)
				queue = append(queue, newPath)
			}
		}
	}
	return nil
}
