deps:
	go mod vendor

## Test stuff
get_lint_config:
	@[ -f ./.golangci.yml ] && echo ".golangci.yml exists" || ( echo "getting .golangci.yml" && curl -O https://raw.githubusercontent.com/microparts/docker-golang/master/.golangci.yml )
.PHONY: get_lint_config

lint: get_lint_config
	golangci-lint run -v
.PHONY: lint

test-unit:
	go test ./... --race --cover -count=1 -timeout 1s -coverprofile=c.out -v
.PHONY: test-unit

coverage-html:
	go tool cover -html=c.out -o coverage.html
.PHONE: coverage-html

test: deps test-unit coverage-html
.PHONY: test
