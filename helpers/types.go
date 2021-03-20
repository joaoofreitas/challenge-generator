package helpers

type Exercise struct {
	name        string
	description string
}

type Subject struct {
	name      string
	exercises []Exercise
}
