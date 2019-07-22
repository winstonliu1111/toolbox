PKGS := $(go list ./... | grep -v /vendor)

BIN_DIR := $(GOPATH)/bin
DEP := $(BIN_DIR)/dep

BINARY := toolbox
VERSION ?= vlatest
PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: test windows linux darwin release vendor clean

release: $(PLATFORMS)
	cd release && shasum -a256 $(BINARY)-* > SHA256SUMS

$(PLATFORMS): vendor
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64

clean:
	rm -rf release

test: vendor
	go test $(PKGS)

vendor: $(DEP)
	dep ensure -v

$(DEP):
	go get -u github.com/golang/dep/cmd/dep
