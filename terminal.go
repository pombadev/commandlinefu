package main

import (
	"os"

	"golang.org/x/term"
)

// terminal state is not restored on exit, manually doing it
// https://github.com/c-bata/go-prompt/issues?q=is%3Aissue+stty

// Thanks @WangYihang
// https://github.com/c-bata/go-prompt/issues/233#issuecomment-934395156

var termState *term.State

func saveTermState() {
	oldState, err := term.GetState(int(os.Stdin.Fd()))

	if err != nil {
		return
	}

	termState = oldState
}

func restoreTermState() {
	if termState != nil {
		run(func() error {
			return term.Restore(int(os.Stdin.Fd()), termState)
		})
	}
}

func init() {
	saveTermState()
}
