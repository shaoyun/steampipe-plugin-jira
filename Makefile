
install:
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/jira@latest/steampipe-plugin-jira.plugin *.go

local:
	go build -o ~/.steampipe/plugins/local/jira/steampipe-plugin-jira.plugin *.go
amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o target/steampipe-plugin-jira.plugin *.go

