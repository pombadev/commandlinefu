<p align="center">
	<a href="https://www.commandlinefu.com" target="_blank">
		<img alt="logo" src="./static/logo.png" />
	</a>
</p>

# commandlinefu [![Git tags](https://img.shields.io/github/v/tag/pjmp/commandlinefu?label=latest%20tag&style=flat)](https://github.com/pjmp/commandlinefu/tags) [![Go Reference](https://pkg.go.dev/badge/github.com/pjmp/commandlinefu.svg)](https://pkg.go.dev/github.com/pjmp/commandlinefu)

# Introduction

> commandlinefu.com is the place to record those command-line gems that you return to again and again. That way others can gain from your CLI wisdom and you from theirs too. All commands can be commented on, discussed and voted up or down.

`commandlinefu` is an unofficial cli/repl client for [commandlinefu.com](https://www.commandlinefu.com), written in [golang](https://golang.org/).

This uses both the provided [official api](https://www.commandlinefu.com/site/api) and scrapes the website because the official api does not provide all features present in the website.

# Demo
[![Demo](./static/demo.svg)](./static/demo.svg)

# Installation

```bash
go get github.com/pjmp/commandlinefu
```

# Usage

```
Usage of commandlinefu:
	-preview-themes
		Preview available themes
	-query string
		Command or question to search
	-repl
		Starts a commandlinefu repl (default true)
	-theme value
		Set syntax highlight theme
	-version
		Prints version information
```
