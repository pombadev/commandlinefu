package main

const (
	AppName    string = "commandlinefu"
	AppVersion string = "v1.3.0"
)

func main() {
	app := NewApp()

	app.Run()
}
