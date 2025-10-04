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

// BFS performs a Breadth-First Search starting from a given node to find the shortest path
// avoids blocked rooms
func BFS(graph *Graph, blocked map[string]bool) []string {
	queue := []string{graph.Start}
	visited := map[string]bool{graph.Start: true}
	prev := map[string]string{}

	for len(queue) > 0 {
		room := queue[0]
		queue = queue[1:]

		if room == graph.End {
			// reconstruct path
			path := []string{}
			for at := graph.End; at != ""; at = prev[at] {
				path = append([]string{at}, path...)
			}
			return path
		}

		for _, neighbor := range graph.Links[room] {
			if !visited[neighbor] && !blocked[neighbor] {
				visited[neighbor] = true
				prev[neighbor] = room
				queue = append(queue, neighbor)
			}
		}
	}
	return nil
}
