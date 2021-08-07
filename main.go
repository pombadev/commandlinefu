package main

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	sliced := strings.Split(d.Text, " ")

	var suggestions []prompt.Suggest

	if len(sliced) > 1 {
		cmd := sliced[0]
		if cmd == "browse" {
			suggestions = []prompt.Suggest{
				{Text: "sort-by-votes", Description: "All-time Greats"},
				{Text: "last-month", Description: "Hot this month"},
				{Text: "last-week", Description: "Weekly"},
				{Text: "last-day", Description: "Daily"},
				{Text: "latest", Description: "Latest"},
			}
		}

	} else {
		suggestions = []prompt.Suggest{
			{Text: "random", Description: "Random"},
			{Text: "forthewicked", Description: "The Wicked"},
			{Text: "browse", Description: "browse by many params"},
			{Text: "match", Description: "match"},
			{Text: "help", Description: "Help"},
			{Text: "search", Description: "Search"},
		}
	}

	return prompt.FilterFuzzy(suggestions, d.GetWordBeforeCursor(), true)
}

// go run -ldflags "-X 'main.version=1'" .
// go build -ldflags "-X 'main.version=1'"
var (
	revision string
	version  string
	AppName  string = "commandlinefu"
)

func main() {
	app := NewApp()

	fmt.Printf("A cli and REPL for %s.com (v%s git+%s)\nPlease use `exit` or `Ctrl-D` to exit this program.\nType help to see all available commands and parameter\n", AppName, version, revision)

	repl := prompt.New(
		func(input string) {
			var (
				cmd   string
				param string
			)

			sliced := strings.Split(input, " ")

			if len(sliced) > 1 {
				cmd = sliced[0]
				param = strings.Join(sliced[1:], " ")
			} else {
				cmd = sliced[0]
			}

			switch cmd {
			case "random":
				run(func() error {
					return app.random()
				})
			case "forthewicked":
				run(func() error {
					return app.wicked()
				})
			case "browse":
				run(func() error {
					return app.browse(param)
				})
			case "match":
				run(func() error {
					return app.matching(param)
				})
			case "search":
				run(func() error {
					return app.search(param)
				})
			case "help":
				help("")

			default:
				help(input)
			}
		},
		completer,
		prompt.OptionTitle(AppName),
		prompt.OptionPrefixTextColor(prompt.DarkGreen),
		prompt.OptionInputTextColor(prompt.Green),
		prompt.OptionSuggestionBGColor(prompt.Green),
		prompt.OptionSuggestionTextColor(prompt.Black),
		prompt.OptionSelectedSuggestionTextColor(prompt.White),
	)

	repl.Run()

}
