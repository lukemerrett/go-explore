package game

import (
	"github.com/lukemerrett/go-explore/model"
	"testing"
)

func TestGetFirstScene_GetsFirstScene(t *testing.T) {
	gameData := model.GameData{
		Scenes: map[string]model.Scene{
			"Home": model.Scene{
				Title: "Next scene",
			},
			"Garden": model.Scene{
				Title:      "First scene",
				StartScene: true,
			},
		},
	}

	scene := getFirstScene(gameData)

	if scene.Title != "First scene" {
		t.Fail()
	}
}
