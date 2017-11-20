package model

// GameData holds a list of the scenes available within a game
type GameData struct {
	Scenes map[string]Scene `yaml:"Scenes"`
}

// Scene holds the information to show the user about a particular scene and what they can do next
type Scene struct {
	Title       string            `yaml:"Title"`
	Body        string            `yaml:"Body"`
	Transitions map[string]string `yaml:"Transitions"`
}
