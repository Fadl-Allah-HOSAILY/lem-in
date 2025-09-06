package main

type Graph struct {
	Rooms map[string]Room
	Links map[string][]string
	Start string
	End   string
}
