package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X, Y int
}

type Link struct {
	From string
	To   string
}

// ParseFile reads the lem-in input file and returns the ant count, rooms, links, and any error.
func ParseFile(filename string) (Ants int, Rooms map[string]Room, M_links map[string][]string, StartRoom, EndRoom, Mouvement string, err error) {
	Rooms = make(map[string]Room)
	M_links = make(map[string][]string)
	S_Links := []Link{}

	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, nil, "", "", "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nextStart, nextEnd := false, false
	graphDone := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Comments
		if strings.HasPrefix(line, "#") {
			if line == "##start" {
				nextStart = true
			} else if line == "##end" {
				nextEnd = true
			}
			continue
		}

		// --- Phase 1: Graph parsing ---
		if !graphDone {
			// Ant count
			if Ants == 0 {
				n, parseErr := strconv.Atoi(line)
				if parseErr != nil || n < 1 {
					return 0, nil, nil, "", "", "", fmt.Errorf("invalid ant number: %s", line)
				}
				Ants = n
				continue
			}

			parts := strings.Fields(line)

			// Room definition
			if len(parts) == 3 {
				name := parts[0]
				if strings.HasPrefix(name, "L") || strings.Contains(name, " ") || name == "" {
					return 0, nil, nil, "", "", "", fmt.Errorf("invalid room name: %s", name)
				}
				x, err1 := strconv.Atoi(parts[1])
				y, err2 := strconv.Atoi(parts[2])
				if err1 != nil || err2 != nil {
					return 0, nil, nil, "", "", "", fmt.Errorf("invalid room coordinates: %s", line)
				}
				if _, exists := Rooms[name]; exists {
					return 0, nil, nil, "", "", "", fmt.Errorf("duplicate room: %s", name)
				}

				Rooms[name] = Room{Name: name, X: x, Y: y}
				if nextStart {
					StartRoom, nextStart = name, false
				}
				if nextEnd {
					EndRoom, nextEnd = name, false
				}
				continue
			}

			// Link definition
			if strings.Contains(line, "-") {
				linkParts := strings.Split(line, "-")
				if len(linkParts) != 2 {
					return 0, nil, nil, "", "", "", fmt.Errorf("invalid link format: %s", line)
				}
				a, b := linkParts[0], linkParts[1]
				if _, ok := Rooms[a]; !ok {
					return 0, nil, nil, "", "", "", fmt.Errorf("unknown room in link: %s", a)
				}
				if _, ok := Rooms[b]; !ok {
					return 0, nil, nil, "", "", "", fmt.Errorf("unknown room in link: %s", b)
				}
				S_Links = append(S_Links, Link{From: a, To: b})
				continue
			}

			// If we encounter a line starting with "L", switch to movement phase
			if strings.HasPrefix(line, "L") {
				graphDone = true
			} else {
				return 0, nil, nil, "", "", "", fmt.Errorf("invalid line in graph definition: %s", line)
			}
		}

		// --- Phase 2: Movements parsing ---
		if graphDone {
			if !strings.HasPrefix(line, "L") {
				return 0, nil, nil, "", "", "", fmt.Errorf("invalid movement line: %s", line)
			}

			fields := strings.Fields(line)
			for _, f := range fields {
				i := strings.Index(f, "-")
				if i == -1 {
					return 0, nil, nil, "", "", "", fmt.Errorf("invalid movement format: %s", f)
				}
				roomName := f[i+1:]
				if _, ok := Rooms[roomName]; !ok {
					return 0, nil, nil, "", "", "", fmt.Errorf("unknown room in movement: %s", roomName)
				}
			}
			Mouvement += line + "\n"
		}
	}

	if err = scanner.Err(); err != nil {
		return 0, nil, nil, "", "", "", err
	}

	if Ants == 0 || StartRoom == "" || EndRoom == "" {
		return 0, nil, nil, "", "", "", fmt.Errorf("missing ants or start/end room")
	}

	for _, link := range S_Links {
		M_links[link.From] = append(M_links[link.From], link.To)
		M_links[link.To] = append(M_links[link.To], link.From)
	}

	return Ants, Rooms, M_links, StartRoom, EndRoom, Mouvement, nil
}
