package format

import (
	"github.com/lukemerrett/go-explore/model"
)

// Formatter takes a scene and formats it for user output
type Formatter interface {
	FormatScene(model.Scene) string
	GetNextScene(model.GameData, model.Scene) (model.Scene, error)
	EndingScene()
}
