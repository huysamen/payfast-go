package copyutils

func CopyString(from map[string]any, key string) string {
	if from[key] == nil {
		return ""
	}

	return from[key].(string)
}

func CopyInt(from map[string]any, key string) int {
	if from[key] == nil {
		return 0
	}

	return int(from[key].(float64))
}
