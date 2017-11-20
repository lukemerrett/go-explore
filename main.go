package main

import (
	"fmt"
	"github.com/lukemerrett/go-explore/config"
	"github.com/lukemerrett/go-explore/game"
)

const (
	configFile string = "resources/game-data.yml"
)

func main() {
	gameData, err := config.LoadFromYaml(configFile)
	handleError(err)

	err = game.RunGame(gameData)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
}
