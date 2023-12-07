package main

func cutString(str string, length int) string {
	runes := []rune(str)
	if len(runes) > length {
		return string(runes[:length]) + "..."
	} else {
		return str
	}
}
