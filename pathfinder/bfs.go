package pathfinder

import "lemin/farm"

// ----- Find all shortest paths for each neighbor (without blocking) -----
func FindAllShortestPaths(f farm.Farm) [][]string {
	var allPaths [][]string

	for _, neighbor := range f.Rooms[f.Start].Links {
		// Use BFS to find shortest path for this neighbor
		queue := [][]string{{f.Start, neighbor}}
		visited := make(map[string]bool)
		visited[f.Start] = true
		visited[neighbor] = true

		var shortestPath []string

		for len(queue) > 0 && shortestPath == nil {
			path := queue[0]
			queue = queue[1:]
			current := path[len(path)-1]

			if current == f.End {
				shortestPath = path
				break
			}

			for _, next := range f.Rooms[current].Links {
				if !visited[next] {
					visited[next] = true
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, next)
					queue = append(queue, newPath)
				}
			}
		}

		if shortestPath != nil {
			allPaths = append(allPaths, shortestPath)
		}
	}

	return allPaths
}
