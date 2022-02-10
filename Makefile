SHELL = /usr/bin/bash

clean:
	go clean

build:
	go build

sbuild:
	go build -ldflags "-s -w"

run:
	go run .

.ONESHELL:
tag:
	@force=''
	@tag=$$(grep -w AppVersion main.go | cut -d' ' -f4 | sed 's/"//g')

	if git rev-parse -q --verify "refs/tags/$$tag" &> /dev/null; then
		@echo "'$$tag' exits, overwrite?"
		select yn in "Yes" "No"; do
			case $$yn in
				Yes ) force="-f"; break;;
				No ) exit 0;;
			esac
		done
	fi

	git tag -a "$$tag" -m "Tagging $$tag" $$force
