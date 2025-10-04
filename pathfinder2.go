package main

// ----- Find all shortest paths for each neighbor (without blocking) -----
func findAllShortestPaths(f *Farm) [][]string {
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
// ----- Check if two paths share intermediate rooms -----
func pathsShareRooms(path1, path2 []string) bool {
	// Create set of intermediate rooms for path1
	rooms1 := make(map[string]bool)
	for i := 1; i < len(path1)-1; i++ {
		rooms1[path1[i]] = true
	}

	// Check if path2 uses any of these rooms
	for i := 1; i < len(path2)-1; i++ {
		if rooms1[path2[i]] {
			return true
		}
	}
	return false
}

// ----- Select best non-conflicting paths -----
func selectBestPaths(f *Farm, allPaths [][]string) [][]string {
	if len(allPaths) == 0 {
		return nil
	}

	// Sort paths by length (shortest first)
	for i := 0; i < len(allPaths)-1; i++ {
		for j := i + 1; j < len(allPaths); j++ {
			if len(allPaths[j]) < len(allPaths[i]) {
				allPaths[i], allPaths[j] = allPaths[j], allPaths[i]
			}
		}
	}

	var selected [][]string
	usedRooms := make(map[string]bool)

	for _, path := range allPaths {
		conflict := false

		// Check if this path conflicts with any selected path
		for _, selectedPath := range selected {
			if pathsShareRooms(path, selectedPath) {
				conflict = true
				break
			}
		}

		if !conflict {
			selected = append(selected, path)
			// Mark intermediate rooms as used
			for i := 1; i < len(path)-1; i++ {
				usedRooms[path[i]] = true
			}
		}
	}

	return selected
}
