package utils

import (
	"log"
	"os"
	"runtime"
	"strings"
)

func ReadFileAndGetItsLines(path string) ([]string, error) {
	log.Printf("About to open file: %s", path)
	content, err := os.ReadFile(path)
	if err != nil {
		log.Println("Unable to read file due to: " + err.Error())
		return nil, err
	}
	log.Println("Reading file was successful, about to split it into lines.")
	lines := strings.Split(string(content), GetLineSeparator())
	return lines, nil
}

func GetLineSeparator() string {
	switch currentOs := runtime.GOOS; currentOs {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}
