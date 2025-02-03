package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

const asciiAppLogo = `┌─────────────────────────────────────────────────────────────┐
│                                                             │
│     _____                     _____         _               │
│    |   __|___ ___ ___ ___ ___|     |___ ___| |_ ___ _ _     │
│    |__   |  _|  _| .'| . | -_| | | | . |   | '_| -_| | |    │
│    |_____|___|_| |__,|  _|___|_|_|_|___|_|_|_,_|___|_  |    │
│                      |_|                           |___|    │
│                                                             │
└─────────────────────────────────────────────────────────────┘`

type menuItem struct {
	text   string
	key    string
	action func()
}

func exit() {
	os.Exit(0)
}

var exitMenuItem = menuItem{text: "Close", key: "esc", action: exit}

var specialKeyMap = map[string]byte{
	"esc":   27,
	"enter": 13,
}

const ctrlCCode = 3

func clearLines(lines int) {
	fmt.Printf("\033[%dA", lines) // Move cursor up
	fmt.Print("\033[0J")          // Clear from cursor to end of screen
}

func handleTerminalMenu(options []menuItem) error {
	text := strings.Builder{}
	for _, option := range options {
		text.WriteString(fmt.Sprintf("[%s] %s ", option.key, option.text))
	}
	text.WriteString("\n")

	fmt.Println("Choose an option:")
	fmt.Println(text.String())

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	for {
		buf := make([]byte, 1)
		_, err = os.Stdin.Read(buf)
		if err != nil {
			return err
		}

		if buf[0] == ctrlCCode {
			exit()
		}

		var action *func()
		for _, option := range options {
			if specialKeyMap[option.key] == buf[0] || strings.EqualFold(option.key, string(buf)) {
				action = &option.action
				break
			}
		}

		if action == nil {
			continue
		}

		clearLines(3)

		(*action)()
	}
}
