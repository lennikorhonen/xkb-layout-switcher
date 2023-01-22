package main

import (
	"log"
	"os/exec"
    "strings"
)

func main() {
    current_str := getCurrentLayout()
    var lang string

    if (current_str == "fi") {
        lang = "us"
    } else if (current_str == "us"){
        lang = "fi"
    } else {
        log.Println("Current not in us or fi. Defaulting to us")
        lang = "us"
    }

    changeCurrentLayout(lang)
}

func getCurrentLayout() string {
    current, err := exec.Command("bash", "-c", `setxkbmap -query | grep -oP 'layout:\s*\K([\w,]+)'`).Output()

    if err != nil {
        log.Println(err)
    }

    current_str := strings.Trim(string(current), "\n")

    return current_str
}

func changeCurrentLayout(lang string) {
    _, err := exec.Command("setxkbmap", lang).Output()

    if err != nil {
        log.Println(err)
    }
}
