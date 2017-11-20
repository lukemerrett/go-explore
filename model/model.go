package model

// GameData holds a list of the scenes available within a game
type GameData struct {
	Scenes map[string]Scene
}

// Scene holds the information to show the user about a particular scene and what they can do next
type Scene struct {
	Title       string
	Body        string
	Transitions []Transition
}

// Transition maps between an option a user can choose and the next scene they will see
type Transition struct {
	Text     string
	SceneKey string
}
