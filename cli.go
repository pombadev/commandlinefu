package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/alecthomas/chroma/styles"

	"github.com/c-bata/go-prompt"
	colour "github.com/fatih/color"
)

var availableStyles = styles.Names()

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

		if cmd == "settheme" {
			for _, theme := range availableStyles {
				suggestions = append(suggestions, prompt.Suggest{
					Text:        theme,
					Description: theme,
				})
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
			{Text: "settheme", Description: "Set syntax highlight theme"},
			{Text: "version", Description: "Prints version information"},
		}
	}

	return prompt.FilterFuzzy(suggestions, d.GetWordBeforeCursor(), true)
}

// Cli Represent our cli as a struct
type Cli struct {
	// Start a REPL session
	repl bool
	// Query string if repl is not started
	query string
	// App's version
	version bool
	// List available themes
	listThemes bool
	// Specify theme to use
	theme string
}

// NewCli Initialize a new instance of `Cli`
func NewCli() Cli {
	repl := flag.Bool("repl", true, fmt.Sprintf("Starts a %s repl", AppName))
	query := flag.String("query", "", "Command or question to search")
	version := flag.Bool("version", false, "Prints version information")
	listThemes := flag.Bool("list-themes", false, "List available themes")

	var theme string = "dracula"

	flag.Func("theme", "Set syntax highlight theme", func(q string) error {
		hasTheme, err := HasTheme(q)

		if hasTheme {
			theme = q
		} else {
			return err
		}

		return nil
	})

	flag.Parse()

	return Cli{repl: *repl, query: *query, version: *version, theme: theme, listThemes: *listThemes}
}

func HasTheme(name string) (bool, error) {
	for _, styleName := range availableStyles {
		if styleName == name {
			return true, nil
		}
	}

	return false, fmt.Errorf("\nValue must be one of\n%s\n", strings.Join(availableStyles, ", "))
}

// Version Show App's version
func (app *App) Version() {
	fmt.Println(AppName + " " + AppVersion)
}

// Repl Start a new REPL session
func (app *App) Repl() {
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
			case "settheme":
				hasTheme, err := HasTheme(param)
				if hasTheme {
					app.cli.theme = param
				} else {
					fmt.Print(err)
				}
			case "version":
				app.Version()
			case "exit":
				os.Exit(0)
			case "help":
				help("")
			default:
				help(input)
			}
		},
		completer,
		prompt.OptionMaxSuggestion(20),
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
func (app *App) Search() {
	run(func() error {
		return app.search(app.cli.query)
	})
}

// ListThemes List available themes
func (app *App) ListThemes() {
	source := `#!/usr/bin/env sh

# All fits on one line
command1 | command2

# Long commands
command1 \
  | command2 \
  | command3 \
  | command4

# log to stderr
err() {
  echo "[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $*" >&2
}

if ! do_something; then
  err "Unable to do_something"
  exit 1
fi

`

	cl := colour.New(colour.FgWhite).Add(colour.Underline).Add(colour.Bold)
	length := len(availableStyles) - 1

	for index, style := range availableStyles {
		cl.Printf("[%d/%d] %s\n\n", index, length, style)
		quick.Highlight(os.Stdout, source, "bash", "terminal256", style)
	}
}
