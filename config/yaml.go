package config

import (
	"github.com/go-yaml/yaml"
	"github.com/lukemerrett/go-explore/model"
	"io/ioutil"
)

// LoadFromYaml takes a filename and loads the game data from a yaml file
func LoadFromYaml(fileName string) (model.GameData, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return model.GameData{}, err
	}

	gameData := model.GameData{}
	err = yaml.Unmarshal(bytes, &gameData)
	return gameData, err
}
