package format

import (
	"bytes"
	"fmt"
	"github.com/lukemerrett/go-explore/model"
)

// ConsoleFormatter takes a scene and formats it for display in the console
type ConsoleFormatter struct {
}

// FormatScene takes a scene and formats it for display in the console
func (ConsoleFormatter) FormatScene(scene model.Scene) string {
	var buffer bytes.Buffer
	buffer.WriteString("\n\n---------------------------------\n")
	buffer.WriteString(fmt.Sprintf("%s\n\n%s", scene.Title, scene.Body))

	if transitionPresent(scene.Transitions) {
		buffer.WriteString(fmt.Sprintf("\n\nOptions:"))
		i := 1
		for _, value := range scene.Transitions {
			buffer.WriteString(fmt.Sprintf("\n%v. %s", i, value))
			i++
		}
	}

	return buffer.String()
}
