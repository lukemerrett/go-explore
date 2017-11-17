package model

// Scene holds the information to show the user about a particular scene and what they can do next
type Scene struct {
	Title   string
	Body    string
	Options []Trigger
}

// Trigger maps between an option a user can choose and the next scene they will see
type Trigger struct {
	Option string
	Scene  string
}
