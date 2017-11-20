package config

import (
	"fmt"
	"testing"
)

func TestLoadFromYaml_CanLoadYamlFile(t *testing.T) {
	filename := "test-data.yml"
	gameData, err := LoadFromYaml(filename)

	if err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}

	scene := gameData.Scenes["Home"]

	assert(t, "Your home", scene.Title)
	assert(t, "A house on a street in a neighbourhood", scene.Body)
	assert(t, "Go to the garden", scene.Transitions["Garden"])
	assert(t, "Go inside the house", scene.Transitions["Hallway"])
}

func TestLoadFromYaml_FileDoesntExist(t *testing.T) {
	filename := "asdqewrqfeesf.s"
	_, err := LoadFromYaml(filename)
	if err == nil {
		t.Fail()
	}
}

func assert(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Fail()
	}
}
