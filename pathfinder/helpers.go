package pathfinder

import "lemin/farm"

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
func SelectBestPaths(f farm.Farm, allPaths [][]string) [][]string {
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

func FindNonOverlappingPaths(f farm.Farm) [][]string {
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
func bfsShortestPath(f farm.Farm, startNeighbor string, blockedRooms map[string]bool) []string {
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
