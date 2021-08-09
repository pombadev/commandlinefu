package main

const (
	AppName    string = "commandlinefu"
	AppVersion string = "v2.0.1"
)

func main() {
	app := NewApp()

	app.Run()
}
