package main

type Graph struct {
	Rooms map[string]Room
	Links map[string][]string
	Start string
	End   string
}

// FindMultiplePaths finds all disjoint paths from start to end
func FindMultiplePaths(graph *Graph) []Path {
	paths := []Path{}
	blocked := map[string]bool{}

	for {
		rawPath := BFS(graph, blocked)
		if rawPath == nil {
			break
		}

		paths = append(paths, Path{Rooms: rawPath, Length: len(rawPath)})

		// block intermediate rooms (except start/end)
		for _, room := range rawPath {
			if room != graph.Start && room != graph.End {
				blocked[room] = true
			}
		}
	}

	return paths
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
