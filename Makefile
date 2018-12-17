GO := go
GOARCH=amd64
GOOS=$(shell uname -s | tr '[:upper:]' '[:lower:]')
GOPATH ?= $(GOPATH:)

PLUGINNAME=boilerplate
PLUGINVERSION=$(shell echo `cat plugin.json | grep -Po '"'"version"'"\s*:\s*"\K([^"]*)'`)
PACKAGENAME=mattermost-plugin-$(PLUGINNAME)-v$(PLUGINVERSION)

BLACK=`tput setaf 0`
RED=`tput setaf 1`
GREEN=`tput setaf 2`
YELLOW=`tput setaf 3`
BLUE=`tput setaf 4`
MAGENTA=`tput setaf 5`
CYAN=`tput setaf 6`
WHITE=`tput setaf 7`

BOLD=`tput bold`
INVERSE=`tput rev`
RESET=`tput sgr0`

.PHONY: default test clean check-style checkjs checkgo govet golint gofmt .distclean dist fix fixjs fixgo

default: check-style test dist

check-style: .npminstall vendor
	@echo ${BOLD}"Checking for style guide compliance\n"${RESET}
	@make checkjs
	@make checkgo

checkjs: webapp
	@echo ${BOLD}Running ESLINT${RESET}
	@cd webapp && npm run lint
	@echo ${GREEN}"eslint success\n"${RESET}

checkgo: server govet golint gofmt

govet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		echo "--> installing govet"; \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@echo ${BOLD}Running GOVET${RESET}${RED}
	@cd server
	$(eval PKGS := $(shell go list ./... | grep -v /vendor/))
	@$(GO) vet $(PKGS)
	@echo ${GREEN}"govet success\n"${RESET}

golint:
	@command -v golint >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing golint"; \
		go get -u golang.org/x/lint/golint; \
	fi
	@echo ${BOLD}Running GOLINT${RESET}${RED}
	@cd server
	$(eval PKGS := $(shell go list ./... | grep -v /vendor/))
	@# grep -v is used to ignore specific linting rules and not echo their failure terminal
	-@for pkg in $(PKGS) ; do \
		golint $$pkg | grep -v "have comment" | grep -v "comment on exported" | grep -v "lint suggestions" ; \
	done && echo ${RED}"golint failure\n"${RESET} || echo ${GREEN}"golint success\n"${RESET}

fix: fixjs fixgo

fixjs:
	@echo ${BOLD}Formatting js giles${RESET}
	@cd webapp && npm run fix
	@echo "formatted js files\n"

fixgo:
	@command -v goimports >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing goimports"; \
		go get golang.org/x/tools/cmd/goimports; \
	fi
	@echo ${BOLD}Formatting go giles${RESET}
	@cd server
	@find ./ -type f -name "*.go" -not -path "./server/vendor/*" -exec goimports -w {} \;
	@echo "formatted go files\n"

gofmt:
	@echo ${BOLD}Running GOFMT${RESET}${RED}
	@for package in $$(go list ./server/...); do \
		files=$$(go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' $$package); \
		if [ "$$files" ]; then \
			gofmt_output=$$(gofmt -d -s $$files 2>&1); \
			if [ "$$gofmt_output" ]; then \
				echo "$$gofmt_output"; \
				echo "gofmt failure\n"; \
				exit 1; \
			fi; \
		fi; \
	done
	@echo ${GREEN}"gofmt success\n"${RESET}

test:
	go test -v -coverprofile=coverage.txt ./...

.npminstall: webapp/package-lock.json
	@echo ${BOLD}"Getting dependencies using npm\n"${RESET}
	cd webapp && npm install
	@echo "\n"

vendor: server/glide.lock
	@echo ${BOLD}"Getting dependencies using glide\n"${RESET}
	cd server && go get github.com/Masterminds/glide
	cd server && $(shell go env GOPATH)/bin/glide install
	@echo "\n"

dist: .distclean check-style plugin.json
	@echo ${BOLD}"Building plugin\n"${RESET}

	# Build and copy files from webapp
	cd webapp && npm run build
	mkdir -p dist/$(PLUGINNAME)/webapp
	cp -r webapp/dist/* dist/$(PLUGINNAME)/webapp/

	# Build files from server
	cd server && go get github.com/mitchellh/gox
	$(shell go env GOPATH)/bin/gox -osarch='darwin/amd64 linux/amd64 windows/amd64' -output 'dist/intermediate/plugin_{{.OS}}_{{.Arch}}' ./server

	# Copy plugin files
	cp plugin.json dist/$(PLUGINNAME)/

	# Copy server executables & compress plugin
	mkdir -p dist/$(PLUGINNAME)/server

	mv dist/intermediate/plugin_darwin_amd64 dist/$(PLUGINNAME)/server/plugin.exe
	cd dist && tar -zcvf $(PACKAGENAME)-darwin-amd64.tar.gz $(PLUGINNAME)/*

	mv dist/intermediate/plugin_linux_amd64 dist/$(PLUGINNAME)/server/plugin.exe
	cd dist && tar -zcvf $(PACKAGENAME)-linux-amd64.tar.gz $(PLUGINNAME)/*

	mv dist/intermediate/plugin_windows_amd64.exe dist/$(PLUGINNAME)/server/plugin.exe
	cd dist && tar -zcvf $(PACKAGENAME)-windows-amd64.tar.gz $(PLUGINNAME)/*

	# Clean up temp files
	rm -rf dist/$(PLUGINNAME)
	rm -rf dist/intermediate

	@echo Linux plugin built at: dist/$(PACKAGENAME)-linux-amd64.tar.gz
	@echo MacOS X plugin built at: dist/$(PACKAGENAME)-darwin-amd64.tar.gz
	@echo Windows plugin built at: dist/$(PACKAGENAME)-windows-amd64.tar.gz

.distclean:
	@echo ${BOLD}"Cleaning dist files\n"${RESET}
	rm -rf dist
	rm -rf webapp/dist
	rm -f server/plugin.exe
	@echo "\n"

clean: .distclean
	@echo ${BOLD}"Cleaning plugin\n"${RESET}
	rm -rf server/vendor
	rm -rf webapp/node_modules
	rm -rf webapp/.npminstall
	@echo "\n"
