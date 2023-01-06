package main

import (
	"log"
	"os/exec"
    "strings"
)

func main() {
    current, err := exec.Command("bash", "-c", "setxkbmap -print | awk '{ print $4 }' | grep -o -e us -e fi").Output()

    if err != nil {
        log.Println(err)
    }
    current_str := strings.Trim(string(current), "\n")

    var lang string

    if (current_str == "fi") {
        lang = "us"
    } else if (current_str == "us"){
        lang = "fi"
    } else {
        log.Println("Current not in us or fi. Defaulting to us")
        lang = "us"
    }

    _, err = exec.Command("setxkbmap", lang).Output()

    if err != nil {
        log.Println(err)
    }
}
