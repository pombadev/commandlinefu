package main

import (
	"flag"
	"fmt"
	"os"
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
			{Text: "exit", Description: "Exit repl session"},
		}
	}

	return prompt.FilterFuzzy(suggestions, d.GetWordBeforeCursor(), true)
}

type Cli struct {
	repl    bool
	query   string
	version bool
	app     App
}

func NewCli() Cli {
	repl := flag.Bool("repl", true, fmt.Sprintf("Starts a %s repl", AppName))
	query := flag.String("query", "", "A query")
	version := flag.Bool("version", false, "Prints version information")

	flag.Parse()

	return Cli{repl: *repl, query: *query, version: *version, app: NewApp()}
}

func (c Cli) Version() {
	fmt.Printf("%s+%s\n", AppVersion, AppRevision)
}

func (c Cli) Repl() {
	fmt.Printf("A cli and REPL for %s.com (%s git+%s)\nPlease use `exit` or `Ctrl-D` to exit this program\nType help to see all available commands and parameter\n", AppName, AppVersion, AppRevision)
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
					return c.app.random()
				})
			case "forthewicked":
				run(func() error {
					return c.app.wicked()
				})
			case "browse":
				run(func() error {
					return c.app.browse(param)
				})
			case "match":
				run(func() error {
					return c.app.matching(param)
				})
			case "search":
				run(func() error {
					return c.app.search(param)
				})
			case "exit":
				os.Exit(0)
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

func (c Cli) Search() {
	run(func() error {
		return c.app.search(c.query)
	})
}
