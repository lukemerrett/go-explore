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

	if optionPresent(scene.Options) {
		buffer.WriteString(fmt.Sprintf("\n\nOptions:"))
		for i, option := range scene.Options {
			buffer.WriteString(fmt.Sprintf("\n%v. %s", i+1, option.Option))
		}
	}
	return buffer.String()
}

// GetNextScene waits for the user to select the next scene from the options
func GetNextScene(currentScene model.Scene) model.Scene {
	optionCount := int64(0)
	if optionPresent(currentScene.Options) {
		optionCount = int64(len(currentScene.Options))
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n\nPlease choose an option (1-%v):\n", optionCount)

	selectedOption := 1

	for true {
		text, _ := reader.ReadString('\n')
		selectedOption, err := strconv.ParseInt(text, 10, 64)

		if err != nil || selectedOption > optionCount || selectedOption < optionCount {
			fmt.Printf("\n\nPlease enter a number between 1-%v:\n", optionCount)
		} else {
			break
		}
	}

	return currentScene.Options[selectedOption-1].Scene
}

func optionPresent(list []model.Trigger) bool {
	return list != nil && len(list) > 0
}
