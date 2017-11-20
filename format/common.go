package format

func transitionPresent(transitions map[string]string) bool {
	return transitions != nil && len(transitions) > 0
}
