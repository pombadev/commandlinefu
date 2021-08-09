package main

type App struct {
	urls *Urls
	cli  *Cli
}

func NewApp() *App {
	urls := NewUrls()
	cli := NewCli()

	app := App{&urls, &cli}

	return &app
}
