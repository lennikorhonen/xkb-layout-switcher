package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func main() {
	currentLang := currentLayout()
	langs := settings()

	currentIndex := slices.Index(langs["langs"], currentLang)
	nextIndex := currentIndex + 1

	// If next index would be the same as the length of the langs slice then go
	// back to first language on the list
	if nextIndex == len(langs["langs"]) {
		nextIndex = 0
	}

	setCurrentLayout(langs["langs"][nextIndex])
}

func currentLayout() string {
	current, err := exec.Command("bash", "-c", `setxkbmap -query | grep -oP 'layout:\s*\K([\w,]+)'`).Output()

	if err != nil {
		log.Println(err)
	}

	current_str := strings.Trim(string(current), "\n")

	return current_str
}

func setCurrentLayout(lang string) {
	_, err := exec.Command("setxkbmap", lang).Output()

	if err != nil {
		log.Println(err)
	}
}

func settings() (langs map[string][]string) {
	settings, err := readSettingsFile()
	if err != nil {
		log.Fatal("Error when handling settings: ", err)
	}

	if err := json.Unmarshal(settings, &langs); err != nil {
		log.Fatal("Error when handling setting json: ", err)
	}

	return langs
}

func readSettingsFile() ([]byte, error) {
	dir, dirErr := os.UserConfigDir()
	if dirErr != nil {
		log.Fatal("Failed to read config directory: ", dirErr)
	}

	file := dir + "/kb_layout_switcher/settings.json"
	settings, err := os.ReadFile(file)
	return settings, err
}
