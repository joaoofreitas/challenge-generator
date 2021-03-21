package types

type Exercise struct {
	Name        string
	Description string
}

type Subject struct {
	Name      string
	Exercises []Exercise
}
