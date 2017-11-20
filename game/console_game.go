package game

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/lukemerrett/go-explore/model"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ConsoleController manages the input and output between the game and the user over the command line.
type ConsoleController struct{}

// RunGame runs the whole game loop
func (ConsoleController) RunGame(gameData model.GameData) error {
	currentScene := getFirstScene(gameData)

	for {
		output := formatScene(currentScene)
		fmt.Print(output)

		if currentScene.EndScene {
			endingScene()
			break
		}

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
		if value.StartScene {
			firstScene = value
			break
		}
	}
	return firstScene
}

func formatScene(scene model.Scene) string {
	var buffer bytes.Buffer
	buffer.WriteString("\n\n---------------------------------\n")
	buffer.WriteString(fmt.Sprintf("%s\n\n%s", scene.Title, scene.Body))

	if transitionPresent(scene.Transitions) {
		buffer.WriteString(fmt.Sprintf("\n\nOptions:"))
		i := 1
		for _, key := range orderKeys(scene.Transitions) {
			buffer.WriteString(fmt.Sprintf("\n%v. %s", i, scene.Transitions[key]))
			i++
		}
	}

	return buffer.String()
}

func getNextScene(gameData model.GameData, currentScene model.Scene) (model.Scene, error) {
	transitionCount := 0
	if transitionPresent(currentScene.Transitions) {
		transitionCount = len(currentScene.Transitions)
	}

	reader := bufio.NewReader(os.Stdin)

	switch transitionCount {
	case 1:
		fmt.Printf("\n\nPlease press 1:\n")
	default:
		fmt.Printf("\n\nPlease choose an option (1-%v):\n", transitionCount)
	}

	selectedOption, err := getSelectedOption(reader, transitionCount)
	if err != nil {
		return model.Scene{}, err
	}

	i := 1
	for _, key := range orderKeys(currentScene.Transitions) {
		if i == selectedOption {
			return gameData.Scenes[key], nil
		}
		i++
	}

	return model.Scene{}, fmt.Errorf("Could not match scene")
}

func endingScene() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nPress enter to exit")
	reader.ReadString('\n')
}

func getSelectedOption(reader *bufio.Reader, transitionCount int) (int, error) {
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)

		selectedOption, err := strconv.Atoi(text)

		if err != nil || selectedOption > transitionCount || selectedOption < 1 {
			switch transitionCount {
			case 1:
				fmt.Printf("\n\nPlease press 1:\n")
			default:
				fmt.Printf("\n\nPlease enter a number between 1-%v:\n", transitionCount)
			}
		} else {
			return selectedOption, nil
		}
	}
}

func orderKeys(transitions map[string]string) []string {
	keys := make([]string, 0)
	for key := range transitions {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func transitionPresent(transitions map[string]string) bool {
	return transitions != nil && len(transitions) > 0
}
