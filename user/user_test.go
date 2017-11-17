package user

import (
	"fmt"
	"github.com/lukemerrett/go-explore/model"
	"testing"
)

func TestOutputScene_WithOptions(t *testing.T) {
	scene := model.Scene{
		Title: "Sample",
		Body:  "A body of text",
		Options: []model.Trigger{
			model.Trigger{
				Option: "An option",
				Scene:  model.Scene{},
			},
			model.Trigger{
				Option: "Another option",
				Scene:  model.Scene{},
			},
		},
	}

	output := OutputScene(scene)

	expectedValue := "Sample\n\nA body of text\n\n" +
		"Options:\n1. An option\n2. Another option"

	if expectedValue != output {
		outputExpectedVsActual(expectedValue, output)
		t.Fail()
	}
}

func TestOutputScene_WithNilOptions(t *testing.T) {
	scene := model.Scene{
		Title:   "Sample",
		Body:    "A body of text",
		Options: nil,
	}

	output := OutputScene(scene)

	expectedValue := "Sample\n\nA body of text"

	if expectedValue != output {
		outputExpectedVsActual(expectedValue, output)
		t.Fail()
	}
}

func TestOutputScene_WithEmptyOptions(t *testing.T) {
	scene := model.Scene{
		Title:   "Sample",
		Body:    "A body of text",
		Options: []model.Trigger{},
	}

	output := OutputScene(scene)

	expectedValue := "Sample\n\nA body of text"

	if expectedValue != output {
		outputExpectedVsActual(expectedValue, output)
		t.Fail()
	}
}

func outputExpectedVsActual(expected string, actual string) {
	fmt.Printf("Expected: '%s'\nActual: '%s'", expected, actual)
}
