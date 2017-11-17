package user

import (
	"bytes"
	"fmt"
	"github.com/lukemerrett/go-explore/model"
)

// OutputScene takes a scene and builds the string to show to the user
func OutputScene(scene model.Scene) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s\n\n%s", scene.Title, scene.Body))

	if optionPresent(scene.Options) {
		buffer.WriteString(fmt.Sprintf("\n\nOptions:"))
		for i, option := range scene.Options {
			buffer.WriteString(fmt.Sprintf("%b. %s", i+1, option))
		}
	}
	return buffer.String()
}

func optionPresent(list []model.Trigger) bool {
	return list != nil && len(list) > 0
}
