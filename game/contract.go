package game

import (
	"github.com/lukemerrett/go-explore/model"
)

// Controller is responsible for managing the input and output between the game and the user
type Controller interface {
	RunGame(model.GameData) error
}
