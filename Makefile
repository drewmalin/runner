APP=runner
VERSION=1.0

# Root build binaries
GOCMD=go
DEPCMD=dep

# Build binary subcommands
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt

# Project paths, fields
BINPATH=bin
ROOTPATH=github.com/drewmalin/runner
MAINPATH=$(ROOTPATH)/cmd/runner
VERSIONFIELD=$(ROOTPATH)/internal/config.RunnerVersion

BUILDCMD=$(GOBUILD) -ldflags "-X $(VERSIONFIELD)=$(VERSION)"

build: clean linux darwin

linux:
	env GOOS=linux GOARCH=amd64 $(BUILDCMD) \
		-o $(BINPATH)/nix/$(APP) -v $(MAINPATH)

darwin:
	env GOOS=darwin GOARCH=amd64 $(BUILDCMD) \
		-o $(BINPATH)/osx/$(APP) -v $(MAINPATH)

test_unit:
	$(GOTEST) -v $(ROOTPATH)/cmd/...
	$(GOTEST) -v $(ROOTPATH)/internal/...
	$(GOTEST) -v $(ROOTPATH)/pkg/...

deps:
	$(DEPCMD) ensure

fmt:
	$(GOFMT) fmt ./...

clean:
	$(GOCLEAN)
	rm -fr $(BINPATH)
