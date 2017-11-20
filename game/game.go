package game

import (
	"fmt"
	"github.com/lukemerrett/go-explore/format"
	"github.com/lukemerrett/go-explore/model"
)

// RunGame runs the whole game loop
func RunGame(gameData model.GameData, formatter format.Formatter) error {
	currentScene := getFirstScene(gameData)

	for {
		output := formatter.FormatScene(currentScene)
		fmt.Print(output)

		if currentScene.EndScene {
			break
		}

		nextScene, err := formatter.GetNextScene(gameData, currentScene)
		if err != nil {
			return err
		}
		currentScene = nextScene
	}

	return nil
}

func getFirstScene(gameData model.GameData) model.Scene {
	firstScene := model.Scene{}
	for _, value := range gameData.Scenes {
		if value.StartScene {
			firstScene = value
			break
		}
	}
	return firstScene
}
