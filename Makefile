GOOS=$(shell uname -s | tr '[:upper:]' '[:lower:]')
GOARCH=amd64
PLUGINNAME=boilerplate
PACKAGENAME=mattermost-plugin-$(PLUGINNAME)

.PHONY: default build test run clean stop check-style gofmt .distclean dist

default: check-style test dist

check-style: .npminstall gofmt
	@echo Checking for style guide compliance

	cd webapp && npm run lint

gofmt:
	@echo Running GOFMT

	@for package in $$(go list ./server/...); do \
		echo "Checking "$$package; \
		files=$$(go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' $$package); \
		if [ "$$files" ]; then \
			gofmt_output=$$(gofmt -d -s $$files 2>&1); \
			if [ "$$gofmt_output" ]; then \
				echo "$$gofmt_output"; \
				echo "gofmt failure"; \
				exit 1; \
			fi; \
		fi; \
	done
	@echo "gofmt success"; \

test:
	go test -v -coverprofile=coverage.txt ./...

.npminstall: webapp/package-lock.json
	@echo Getting dependencies using npm

	cd webapp && npm install

vendor: server/glide.lock
	cd server && go get github.com/Masterminds/glide
	cd server && $(shell go env GOPATH)/bin/glide install

dist: .distclean .npminstall vendor plugin.json
	@echo Building plugin

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

run: .npminstall
	@echo Not yet implemented

stop:
	@echo Not yet implemented

.distclean:
	@echo Cleaning dist files

	rm -rf dist
	rm -rf webapp/dist
	rm -f server/plugin.exe

clean: .distclean
	@echo Cleaning plugin

	rm -rf webapp/node_modules
	rm -rf webapp/.npminstall
