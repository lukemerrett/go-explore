package game

import (
	"fmt"
	"github.com/lukemerrett/go-explore/model"
	"testing"
)

func TestOutputScene_WithOptions(t *testing.T) {
	scene := model.Scene{
		Title: "Sample",
		Body:  "A body of text",
		Transitions: map[string]string{
			"Scene 1": "An option",
			"Scene 2": "Another option",
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
		Title:       "Sample",
		Body:        "A body of text",
		Transitions: nil,
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
		Title:       "Sample",
		Body:        "A body of text",
		Transitions: make(map[string]string),
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