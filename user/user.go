package user

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/lukemerrett/go-explore/model"
	"os"
	"strconv"
)

// OutputScene takes a scene and builds the string to show to the user
func OutputScene(scene model.Scene) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s\n\n%s", scene.Title, scene.Body))

	if transitionPresent(scene.Transitions) {
		buffer.WriteString(fmt.Sprintf("\n\nOptions:"))
		for i, transition := range scene.Transitions {
			buffer.WriteString(fmt.Sprintf("\n%v. %s", i+1, transition.Text))
		}
	}
	return buffer.String()
}

// GetNextScene waits for the user to select the next scene from the options
func GetNextScene(gameData model.GameData, currentScene model.Scene) model.Scene {
	transitionCount := int64(0)
	if transitionPresent(currentScene.Transitions) {
		transitionCount = int64(len(currentScene.Transitions))
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n\nPlease choose an option (1-%v):\n", transitionCount)

	selectedOption := 1

	for true {
		text, _ := reader.ReadString('\n')
		selectedOption, err := strconv.ParseInt(text, 10, 64)

		if err != nil || selectedOption > transitionCount || selectedOption < transitionCount {
			fmt.Printf("\n\nPlease enter a number between 1-%v:\n", transitionCount)
		} else {
			break
		}
	}

	sceneKey := currentScene.Transitions[selectedOption-1].SceneKey
	return gameData.Scenes[sceneKey]
}

func transitionPresent(list []model.Transition) bool {
	return list != nil && len(list) > 0
}
