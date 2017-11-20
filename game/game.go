package game

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/lukemerrett/go-explore/model"
	"os"
	"strconv"
)

// RunGame runs the whole game loop
func RunGame(gameData model.GameData) error {
	currentScene := getFirstScene(gameData)

	for true {
		output := outputScene(currentScene)
		fmt.Print(output)
		nextScene, err := getNextScene(gameData, currentScene)
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
		firstScene = value
		break
	}
	return firstScene
}

func outputScene(scene model.Scene) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s\n\n%s", scene.Title, scene.Body))

	if transitionPresent(scene.Transitions) {
		buffer.WriteString(fmt.Sprintf("\n\nOptions:"))
		i := 1
		for _, value := range scene.Transitions {
			buffer.WriteString(fmt.Sprintf("\n%v. %s", i, value))
			i++
		}
	}
	return buffer.String()
}

func getNextScene(gameData model.GameData, currentScene model.Scene) (model.Scene, error) {
	transitionCount := int64(0)
	if transitionPresent(currentScene.Transitions) {
		transitionCount = int64(len(currentScene.Transitions))
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n\nPlease choose an option (1-%v):\n", transitionCount)

	selectedOption := 0

	for true {
		text, _ := reader.ReadString('\n')
		selectedOption, err := strconv.ParseInt(text, 10, 64)

		if err != nil || selectedOption > transitionCount || selectedOption < transitionCount {
			fmt.Printf("\n\nPlease enter a number between 1-%v:\n", transitionCount)
		} else {
			break
		}
	}

	i := 0
	for key := range currentScene.Transitions {
		if i == selectedOption {
			return gameData.Scenes[key], nil
		}
	}

	return model.Scene{}, errors.New("Could not match scene")
}

func transitionPresent(transitions map[string]string) bool {
	return transitions != nil && len(transitions) > 0
}
