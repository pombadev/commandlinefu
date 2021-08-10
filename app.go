package main

import (
	"os"
)

type App struct {
	urls *Urls
	cli  *Cli
}

func (app *App) Run() {
	if app.cli.listThemes {
		app.ListThemes()
		os.Exit(0)
	}

	if app.cli.version {
		app.Version()
		os.Exit(0)
	}

	if len(app.cli.query) >= 1 {
		app.Search()
		os.Exit(0)
	}

	if app.cli.repl {
		app.Repl()
	}
}

func NewApp() *App {
	urls := NewUrls()
	cli := NewCli()

	app := App{&urls, &cli}

	return &app
}
