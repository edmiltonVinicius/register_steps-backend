package utils

import (
	"os"
	"regexp"
)

func GetRootPath() (path string) {
	projectDirName := "register-steps/backend"
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	path = string(rootPath)
	return
}