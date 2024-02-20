package common

import (
	"fmt"
	"strings"
)

func FmtURL(baseURL, port, path string) string {
	if !strings.HasPrefix(baseURL, "http://") || strings.HasPrefix(baseURL, "https://") {
		baseURL = "http://" + baseURL
	}
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}
	if port == "" {
		return fmt.Sprintf("%s/%s", baseURL, path)
	}
	return fmt.Sprintf("%s:%s/%s", baseURL, port, path)
}
