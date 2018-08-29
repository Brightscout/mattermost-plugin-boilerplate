# Mattermost Plugin Boilerplate

#### How to lint Go files ?

1. From server directory, use `goimports -w .` to fix formatting according to gofmt and import ordering.
2. From server directory, use `go list ./... | grep -v /vendor/ | xargs -L1 golint` to lint go files excluding vendor.

#### How to lint JS files ?
1. Go to directory server.
2. Run `npm run lint` in terminal.
