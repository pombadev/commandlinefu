package main

import (
	"io"
	"os/exec"
	"runtime"
)

const (
	AppName    string = "commandlinefu"
	AppVersion string = "v1.5.2"
)

func main() {
	NewApp().Run()

	defer func() {
		// terminal state is not restored on exit, manually doing it
		// https://github.com/c-bata/go-prompt/issues?q=is%3Aissue+stty
		if runtime.GOOS == "linux" {
			cmd := exec.Command("stty", "sane")
			cmd.Stdout = io.Discard
			cmd.Stderr = io.Discard

			_ = cmd.Run()
		}
	}()
}
