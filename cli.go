package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

// Suggestion for available options available within the live REPL session
func completer(d prompt.Document) []prompt.Suggest {
	sliced := strings.Split(d.Text, " ")

	var suggestions []prompt.Suggest

	if len(sliced) > 1 {
		cmd := sliced[0]
		if cmd == "browse" {
			suggestions = []prompt.Suggest{
				{Text: "last-day", Description: "Daily"},
				{Text: "last-month", Description: "Hot this month"},
				{Text: "last-week", Description: "Weekly"},
				{Text: "latest", Description: "Latest"},
				{Text: "sort-by-votes", Description: "All-time Greats"},
			}
		}

	} else {
		suggestions = []prompt.Suggest{
			{Text: "browse", Description: "Browse all commands, sorted by days, month, weekly, all time etc"},
			{Text: "exit", Description: "Exit the current repl session"},
			{Text: "forthewicked", Description: "Commands for the wicked, be warned!"},
			{Text: "help", Description: "Prints help information"},
			{Text: "match", Description: "Match all commands for the given query (searches on comments also)"},
			{Text: "random", Description: "Get random tips"},
			{Text: "search", Description: "Search for commands that matches the given query"},
			{Text: "version", Description: "Prints version information"},
		}
	}

	return prompt.FilterFuzzy(suggestions, d.GetWordBeforeCursor(), true)
}

// Cli Represent our cli a struct
type Cli struct {
	// Start a REPL session?
	repl bool
	// Query string if repl is not started
	query string
	// App's version
	version bool
	// Instance of Commandlinefu struct
	app Commandlinefu
}

// NewCli Initialize a new instance of `Cli`
func NewCli() Cli {
	repl := flag.Bool("repl", true, fmt.Sprintf("Starts a %s repl", AppName))
	query := flag.String("query", "", "Command or question to search")
	version := flag.Bool("version", false, "Prints version information")

	flag.Parse()

	return Cli{repl: *repl, query: *query, version: *version, app: NewCommandlinefu()}
}

// Version Show App's version
func (c Cli) Version() {
	fmt.Println(AppName + " " + AppVersion)
}

// Repl Start a new REPL session
func (c Cli) Repl() {
	var header strings.Builder

	header.WriteString(fmt.Sprintf("A cli and REPL for %s.com (%s)\n", AppName, AppVersion))
	header.WriteString("Please use `exit` or `Ctrl-D` to exit this program\n")
	header.WriteString("Type help to see all available commands and parameter\n")

	fmt.Println(header.String())

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
			case "version":
				c.Version()
			case "exit":
				os.Exit(0)
			case "help":
				help("")
			default:
				help(input)
			}
		},
		completer,
		prompt.OptionMaxSuggestion(uint16(len(completer(prompt.Document{})))),
		prompt.OptionTitle(AppName),
		prompt.OptionPrefixTextColor(prompt.DarkGreen),
		prompt.OptionInputTextColor(prompt.Green),
		prompt.OptionSuggestionBGColor(prompt.Green),
		prompt.OptionSuggestionTextColor(prompt.Black),
		prompt.OptionSelectedSuggestionTextColor(prompt.White),
	)

	repl.Run()
}

// Search whatever query (-query flag) was passed
func (c Cli) Search() {
	run(func() error {
		return c.app.search(c.query)
	})
}
