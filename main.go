package main

const (
	AppName    string = "commandlinefu"
	AppVersion string = "v1.5.0"
)

func main() {
	app := NewApp()

	app.Run()
}
