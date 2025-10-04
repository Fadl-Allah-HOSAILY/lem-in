package farm

type Farm struct {
	Ants  int
	Rooms map[string]Room
	Start string
	End   string
}

type Room struct {
	Name  string
	X, Y  int
	Links []string
}
