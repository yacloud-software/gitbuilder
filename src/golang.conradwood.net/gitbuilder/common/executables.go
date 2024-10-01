package common

import "golang.conradwood.net/go-easyops/utils"

// may return "" if not found, otherwise full path to executable
func FindExecutable(com string) string {
	for _, p := range PATH {
		if utils.FileExists(p + "/" + com) {
			return p + "/" + com
		}
	}
	return ""
}
