PLATFORM := $(shell go run internal/tools/arch/arch.go)
GORELEASER_VERSION := 1.12.3
LINTER_VERSION := 1.50.1
UNAME_S := $(shell uname -s)
UNAME_P := $(shell uname -p)

bin:
	@mkdir -p bin/

bin/golangci-lint: | bin
	curl -fsL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${LINTER_VERSION}
	bin/golangci-lint version

bin/goreleaser: | bin
ifeq ($(UNAME_S),Linux)
	wget https://github.com/goreleaser/goreleaser/releases/download/v${GORELEASER_VERSION}/goreleaser_Linux_x86_64.tar.gz -cqO - | tar -xz -C bin goreleaser
else ifeq ($(UNAME_S),Darwin)
	wget https://github.com/goreleaser/goreleaser/releases/download/v${GORELEASER_VERSION}/goreleaser_Darwin_all.tar.gz -cqO - | tar -xz -C bin goreleaser
else
$(error "unkown OS")
endif
	bin/goreleaser --version

pre-check: bin/golangci-lint
	[[ `(gofmt -l .)`x == x ]] || (echo "go fmt failed" && gofmt -l . && exit 1)
	go vet ./...
	golangci-lint run ./...

test: pre-check
	go clean -testcache
	go test ./...

bin/goimports: | bin
	GOBIN=$(abspath bin) go install golang.org/x/tools/cmd/goimports@latest

bin/gofumpt: | bin
	GOBIN=$(abspath bin) go install mvdan.cc/gofumpt@latest

bin/gci: | bin
	GOBIN=$(abspath bin) go install github.com/daixiang0/gci@latest

fmt: bin/goimports bin/gofumpt bin/gci
	@gofumpt -w -l .
	@bin/goimports -w -l .
	@bin/gci write .

install: dist/*$(PLATFORM)/*
	@echo $^
	@mv $^ $${HOME}/bin

test/compile:
	go test --exec=true ./...

updates:
	@go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all

bin/go-licenses: | bin
	GOBIN=$(abspath bin) go install github.com/google/go-licenses@latest

print-licenses: bin/go-licenses
	@bin/go-licenses csv .

check-licenses: bin/go-licenses
	!(bin/go-licenses csv . | grep -E 'GNU|AGPL|GPL|MPL|CPL|CDDL|EPL|CCBYNC|Facebook|WTFPL')
