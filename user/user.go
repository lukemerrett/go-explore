package user

import (
	"fmt"
	"github.com/lukemerrett/go-explore/model"
)

// OutputScene takes a scene and outputs it to the user's console window
func OutputScene(scene model.Scene) {
	fmt.Printf("%s\n\n", scene.Title)
	fmt.Printf("%s\n\n", scene.Body)

	if optionPresent(scene.Options) {
		fmt.Println("Options:")
		for i, option := range scene.Options {
			fmt.Printf("%b. %s", i+1, option)
		}
	}
}

func optionPresent(list []model.Trigger) bool {
	return list != nil && len(list) > 0
}
