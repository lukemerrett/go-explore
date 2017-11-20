package main

import (
	"fmt"
	"github.com/lukemerrett/go-explore/config"
	"github.com/lukemerrett/go-explore/format"
	"github.com/lukemerrett/go-explore/game"
	"github.com/lukemerrett/go-explore/model"
	"strings"
)

const (
	configFile string = "resources/game-data.yml"
)

func main() {
	gameData, err := config.LoadFromYaml(configFile)
	handleError(err)

	formatter, err := getDependencies(gameData)
	handleError(err)

	err = game.RunGame(gameData, formatter)
	handleError(err)
}

func getDependencies(gameData model.GameData) (format.Formatter, error) {
	outputType := strings.ToLower(gameData.Configuration.OutputType)

	switch outputType {
	case "console":
		return format.ConsoleFormatter{}, nil
	}

	return nil, fmt.Errorf("Could not find matching output type for %v", outputType)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
}
