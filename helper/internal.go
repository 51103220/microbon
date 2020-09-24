package helper

import "strings"

func NormalizePath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return path
}
