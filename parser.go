package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lemin/farm"
)

// ParseFile reads the lem-in input file and returns the ant count, rooms, and start/end rooms.
func ParseFile(filename string) (Ants int, Rooms map[string]farm.Room, StartRoom string, EndRoom string, err error) {
	Rooms = make(map[string]farm.Room)

	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, "", "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nextStart := false
	nextEnd := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			if line == "##start" {
				nextStart = true
			} else if line == "##end" {
				nextEnd = true
			}
			continue
		}

		if Ants == 0 {
			n, parseErr := strconv.Atoi(line)
			if parseErr != nil || n < 1 {
				return 0, nil, "", "", fmt.Errorf("invalid ant number: %s", line)
			}
			Ants = n
			continue
		}

		parts := strings.Fields(line)

		if len(parts) == 3 {
			name := parts[0]
			if strings.HasPrefix(name, "L") || strings.Contains(name, " ") || name == "" {
				return 0, nil, "", "", fmt.Errorf("invalid room name: %s", name)
			}

			x, parseErr1 := strconv.Atoi(parts[1])
			y, parseErr2 := strconv.Atoi(parts[2])
			if parseErr1 != nil || parseErr2 != nil {
				return 0, nil, "", "", fmt.Errorf("invalid room coordinates: %s", line)
			}

			if _, exists := Rooms[name]; exists {
				return 0, nil, "", "", fmt.Errorf("duplicated room name: %s", name)
			}

			for _, r := range Rooms {
				if r.X == x && r.Y == y {
					return 0, nil, "", "", fmt.Errorf("duplicated coordinates: %d %d", x, y)
				}
			}

			Rooms[name] = farm.Room{Name: name, X: x, Y: y, Links: []string{}}

			if nextStart {
				StartRoom = name
				nextStart = false
			}
			if nextEnd {
				EndRoom = name
				nextEnd = false
			}
			continue
		}

		if len(parts) == 1 && strings.Contains(parts[0], "-") {
			linkParts := strings.Split(parts[0], "-")
			if len(linkParts) != 2 {
				return 0, nil, "", "", fmt.Errorf("invalid link format: %s", line)
			}
			a, b := linkParts[0], linkParts[1]

			if a == b {
				return 0, nil, "", "", fmt.Errorf("link cannot connect room to itself: %s", line)
			}

			roomA, okA := Rooms[a]
			roomB, okB := Rooms[b]
			if !okA || !okB {
				return 0, nil, "", "", fmt.Errorf("unknown room in link: %s-%s", a, b)
			}

			roomA.Links = append(roomA.Links, b)
			roomB.Links = append(roomB.Links, a)
			Rooms[a] = roomA
			Rooms[b] = roomB
			continue
		}

		return 0, nil, "", "", fmt.Errorf("invalid line: %s", line)
	}

	if err = scanner.Err(); err != nil {
		return 0, nil, "", "", err
	}

	if Ants == 0 || StartRoom == "" || EndRoom == "" {
		return 0, nil, "", "", fmt.Errorf("missing ants or start/end room")
	}

	return Ants, Rooms, StartRoom, EndRoom, nil
}
