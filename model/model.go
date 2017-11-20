package model

// GameData holds a list of the scenes available within a game
type GameData struct {
	Configuration Configuration    `yaml:"Configuration"`
	Scenes        map[string]Scene `yaml:"Scenes"`
}

// Configuration holds the global config for the application
type Configuration struct {
	OutputType string `yaml:"Output Type"`
}

// Scene holds the information to show the user about a particular scene and what they can do next
type Scene struct {
	Title       string            `yaml:"Title"`
	Body        string            `yaml:"Body"`
	EndScene    bool              `yaml:"End Scene"`
	Transitions map[string]string `yaml:"Transitions"`
}
