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

	assertS(t, "Console", gameData.Configuration.OutputType)

	scene := gameData.Scenes["Home"]

	assertS(t, "Your home", scene.Title)
	assertB(t, false, scene.EndScene)
	assertS(t, "A house on a street in a neighbourhood", scene.Body)
	assertS(t, "Go to the garden", scene.Transitions["Garden"])
	assertS(t, "Go inside the house", scene.Transitions["Hallway"])

	scene = gameData.Scenes["Garden"]
	assertS(t, "Your garden", scene.Title)
	assertS(t, "A garden behind the house", scene.Body)
	assertB(t, true, scene.EndScene)
}

func TestLoadFromYaml_FileDoesntExist(t *testing.T) {
	filename := "asdqewrqfeesf.s"
	_, err := LoadFromYaml(filename)
	if err == nil {
		t.Fail()
	}
}

func assertS(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Fail()
	}
}

func assertB(t *testing.T, expected bool, actual bool) {
	if expected != actual {
		t.Fail()
	}
}
