package data

import "strings"

// GetData ...
func GetData() []string {
	return strings.Split(dictionary, "\n")
}
