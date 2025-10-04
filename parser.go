package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lemin/farm"
)

type Link struct {
	From string
	To   string
}

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
	parsingGraph := true // first part is graph, after that → movements

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// If we hit a line starting with "L", we’re in movements
		if strings.HasPrefix(line, "L") {
			parsingGraph = false
		}

		if parsingGraph {
			if strings.HasPrefix(line, "#") {
				if line == "##start" {
					nextStart = true
				} else if line == "##end" {
					nextEnd = true
				}
				continue
			}

			// Parse ant count
			if Ants == 0 {
				n, parseErr := strconv.Atoi(line)
				if parseErr != nil || n < 1 {
					return 0, nil, "", "", fmt.Errorf("invalid ant number: %s", line)
				}
				Ants = n
				continue
			}

			parts := strings.Fields(line)

			// Room definition
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
					return 0, nil, "", "", fmt.Errorf("duplicated room: %s", name)
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

			// Link definition
			if strings.Contains(line, "-") {
				linkParts := strings.Split(line, "-")
				if len(linkParts) != 2 {
					return 0, nil, "", "", fmt.Errorf("invalid link format: %s", line)
				}
				a, b := linkParts[0], linkParts[1]
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
		} else {
			return 0, nil, "", "", fmt.Errorf("invalid line: %s", line)
		}
	}

	if err = scanner.Err(); err != nil {
		return 0, nil, "", "", err
	}

	if Ants == 0 || StartRoom == "" || EndRoom == "" {
		return 0, nil, "", "", fmt.Errorf("missing ants or start/end room")
	}

	return Ants, Rooms, StartRoom, EndRoom, nil
}
