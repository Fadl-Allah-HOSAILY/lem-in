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
func ParseFile(filename string) (Ants int, Rooms map[string]Room, M_links map[string][]string, StartRoom string, EndRoom string, err error) {
	Rooms = make(map[string]Room)
	M_links = make(map[string][]string)
	S_Links := []Link{}

	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, nil, "", "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nextStart := false
	nextEnd := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
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
				err = fmt.Errorf("invalid ant number: %s", line)
				return 0, nil, nil, "", "", err
			}
			Ants = n
			continue
		}

		parts := strings.Fields(line)

		// Room definition
		if len(parts) == 3 {
			name := parts[0]
			if strings.HasPrefix(name, "L") || strings.Contains(name, " ") || name == "" {
				err = fmt.Errorf("invalid room name: %s", name)
				return 0, nil, nil, "", "", err
			}

			x, parseErr1 := strconv.Atoi(parts[1])
			y, parseErr2 := strconv.Atoi(parts[2])
			if parseErr1 != nil || parseErr2 != nil {
				err = fmt.Errorf("invalid room coordinates: %s", line)
				return 0, nil, nil, "", "", err
			}

			if _, exists := Rooms[name]; exists {
				err = fmt.Errorf("duplicated room: %s", name)
				return 0, nil, nil, "", "", err
			}

			Rooms[name] = Room{Name: name, X: x, Y: y}

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
				err = fmt.Errorf("invalid link format: %s", line)
				return 0, nil, nil, "", "", err
			}
			a, b := linkParts[0], linkParts[1]
			if _, ok := Rooms[a]; !ok {
				err = fmt.Errorf("unknown room in link: %s", a)
				return 0, nil, nil, "", "", err
			}
			if _, ok := Rooms[b]; !ok {
				err = fmt.Errorf("unknown room in link: %s", b)
				return 0, nil, nil, "", "", err
			}
			S_Links = append(S_Links, Link{From: a, To: b})
			continue
		}

		err = fmt.Errorf("invalid line in input: %s", line)
		return 0, nil, nil, "", "", err
	}

	if err = scanner.Err(); err != nil {
		return 0, nil, nil, "", "", err
	}

	if Ants == 0 || StartRoom == "" || EndRoom == "" {
		err = fmt.Errorf("missing ants or start/end room")
		return 0, nil, nil, "", "", err
	}

	for _, link := range S_Links {
		M_links[link.From] = append(M_links[link.From], link.To)
		M_links[link.To] = append(M_links[link.To], link.From)
	}

	return Ants, Rooms, M_links, StartRoom, EndRoom, nil
}
