package main

type Path struct {
	Rooms  []string
	Length int
}

func distributeAnts(paths []Path, numAnts int) []int {
	antsOnPath := make([]int, len(paths))

	for ; numAnts > 0; numAnts-- {
		minIndex := 0
		minScore := paths[0].Length + antsOnPath[0]
		for i, p := range paths {
			score := p.Length + antsOnPath[i]
			if score < minScore {
				minScore = score
				minIndex = i
			}
		}
		antsOnPath[minIndex]++
	}

	return antsOnPath
}
