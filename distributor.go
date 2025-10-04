package main

func distributeAnts(paths [][]string, numAnts int) []int {
	antsOnPath := make([]int, len(paths))

	for ; numAnts > 0; numAnts-- {
		minIndex := 0
		minScore := len(paths[0]) + antsOnPath[0]

		for i, p := range paths {
			score := len(p) + antsOnPath[i]
			if score < minScore {
				minScore = score
				minIndex = i
			}
		}

		antsOnPath[minIndex]++
	}

	return antsOnPath
}
