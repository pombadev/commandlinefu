package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/alecthomas/chroma/quick"
	"github.com/antchfx/htmlquery"
	"github.com/briandowns/spinner"
	"github.com/c-bata/go-prompt"
	"golang.org/x/net/html"
)

type App struct {
	randomUrl string
	wickedUrl string
	browseUrl string
	matchUrl  string
	searchUrl string
}

func (a App) search(query string) error {
	form := url.Values{}

	form.Set("q", query)

	res, err := http.PostForm(a.searchUrl, form)

	if err != nil {
		return err
	}

	node, err := html.Parse(res.Body)

	if err != nil {
		return err
	}

	searchNodes := htmlquery.Find(node, "//ul/li[*]")

	extras := regexp.MustCompile(`\(\d.*\s.*\)`)

	var sb strings.Builder

	for _, node := range searchNodes {
		cmdNode := htmlquery.FindOne(node, "/div[1]")
		descNode := htmlquery.FindOne(node, "/div[2]")

		cmd := htmlquery.InnerText(cmdNode)
		desc := htmlquery.InnerText(descNode)

		sb.WriteString(
			"# " +
				extras.ReplaceAllString(strings.TrimSpace(desc), "") +
				"\n" +
				cmd +
				"\n\n",
		)
	}

	return prettyPrint(sb.String())
}

func (a App) browse(params string) error {
	browseUrl := a.browseUrl

	if len(params) != 0 {
		browseUrl = fmt.Sprintf("%s/%s", browseUrl, params)
	}

	resp, err := fetch(fmt.Sprintf("%s/plaintext", browseUrl))

	if err != nil {
		return err
	}

	return prettyPrint(trimFirstLine(*resp))
}

func (a App) wicked() error {
	resp, err := http.Get(a.wickedUrl)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	out := strings.TrimSpace(trimFirstLine(string(body)))

	return prettyPrint(out + "\n")
}

func (a App) random() error {
	resp, err := http.Get(a.randomUrl)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	htmlNode, err := htmlquery.LoadURL(resp.Request.URL.String())

	if err != nil {
		return err
	}

	node, err := htmlquery.Query(htmlNode, "//*[@id='terminal-display-main']")

	if err != nil {
		return err
	}

	title := htmlquery.InnerText(htmlquery.FindOne(node, "//h1"))
	desc := htmlquery.InnerText(htmlquery.FindOne(node, "//div[1]/span[2]"))

	source := fmt.Sprintf("# %s\n%s\n", strings.TrimSpace(title), strings.TrimSpace(desc))

	return prettyPrint(source)
}

func (a App) matching(query string) error {
	matchUrl := fmt.Sprintf("%s/%s/%s/plaintext/sort-by-votes", a.matchUrl, query, base64.StdEncoding.EncodeToString([]byte(query)))

	resp, err := fetch(matchUrl)

	if err != nil {
		return err
	}

	return prettyPrint(trimFirstLine(*resp))
}

func fetch(url string) (*string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	rect := string(body)

	return &rect, nil
}

func prettyPrint(source string) error {
	err := quick.Highlight(os.Stdout, source, "bash", "terminal256", "dracula")

	if err != nil {
		return err
	}

	return nil
}

func trimFirstLine(s string) string {
	str := strings.Split(s, "\n")

	var sb strings.Builder

	for _, item := range str[2:] {
		sb.WriteString(strings.TrimSpace(item) + "\n")
	}

	return sb.String()

}

func run(callback func() error) {
	spin := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	spin.Start()
	defer spin.Stop()

	err := callback()

	if err != nil {
		log.Fatal(err)
	}
}

func help(arg string) {
	if len(arg) > 0 {
		fmt.Printf("Invalid command: %s\n", arg)
	}

	fmt.Println("List of available commands are: ")

	for _, n := range completer(prompt.Document{}) {
		fmt.Println(n.Text)
	}
}

func NewApp() App {
	var origin, has = os.LookupEnv("COMMANDLINEFU_HOST")

	if !has {
		origin = "https://www.commandlinefu.com"
	}

	baseUrl := fmt.Sprintf("%s/commands", origin)
	randomUrl := fmt.Sprintf("%s/random", baseUrl)
	wickedUrl := fmt.Sprintf("%s/forthewicked/plaintext", baseUrl)
	browseUrl := fmt.Sprintf("%s/browse", baseUrl)
	matchUrl := fmt.Sprintf("%s/matching", baseUrl)
	searchUrl := fmt.Sprintf("%s/search/autocomplete", origin)

	return App{
		randomUrl,
		wickedUrl,
		browseUrl,
		matchUrl,
		searchUrl,
	}
}
