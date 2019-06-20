# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

ifdef GOPATH
GCFLAGS=-trimpath=$(shell printenv GOPATH)/src
else
GCFLAGS=-trimpath=$(shell go env GOPATH)/src
endif

LDFLAGS=-X main.gitBranch=$(shell git branch | grep \* | cut -d ' ' -f2) -X main.gitCommit=$(shell git rev-list -1 HEAD)

ifndef DEST
DEST=bin
endif

.PHONY: help
help:  ## Print the help documentation
	@grep -E '^[a-zA-Z_-\]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#
# Dependencies
#

deps_go:  ## Install Go dependencies
	go get -d -t ./...

.PHONY: deps_go_test
deps_go_test: ## Download Go dependencies for tests
	go get golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow # download shadow
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow # install shadow
	go get -u github.com/kisielk/errcheck # download and install errcheck
	go get -u github.com/client9/misspell/cmd/misspell # download and install misspell
	go get -u github.com/gordonklaus/ineffassign # download and install ineffassign
	go get -u honnef.co/go/tools/cmd/staticcheck # download and instal staticcheck
	go get -u golang.org/x/tools/cmd/goimports # download and install goimports

deps_arm:  ## Install dependencies to cross-compile to ARM
	# ARMv7
	apt-get install -y libc6-armel-cross libc6-dev-armel-cross binutils-arm-linux-gnueabi libncurses5-dev gcc-arm-linux-gnueabi g++-arm-linux-gnueabi
  # ARMv8
	apt-get install gcc-aarch64-linux-gnu g++-aarch64-linux-gnu

deps_gopherjs:  ## Install GopherJS with jsbuiltin
	go get -u github.com/gopherjs/gopherjs
	go get -u -d -tags=js github.com/gopherjs/jsbuiltin

deps_javascript:  ## Install dependencies for JavaScript tests
	npm install .

#
# Go building, formatting, testing, and installing
#

.PHONY: fmt
fmt:  ## Format Go source code
	go fmt $$(go list ./... )

.PHONY: imports
imports: ## Update imports in Go source code
	goimports -w $$(find . -iname '*.go')

.PHONY: vet
vet: ## Vet Go source code
	go vet $$(go list ./... )

.PHONY: test_go
test_go: ## Run Go tests
	bash scripts/test.sh

build: build_cli build_javascript  ## Build CLI and JavaScript

install:  ## Install gotmpl CLI on current platform
	go install -gcflags="$(GCFLAGS)" -ldflags="$(LDFLAGS)" github.com/spatialcurrent/gotmpl/cmd/gotmpl

#
# Command line Programs
#

bin/gotmpl_darwin_amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/gotmpl_darwin_amd64 -gcflags="$(GCFLAGS)" -ldflags="$(LDFLAGS)" github.com/spatialcurrent/gotmpl/cmd/gotmpl

bin/gotmpl_linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/gotmpl_linux_amd64 -gcflags="$(GCFLAGS)" -ldflags="$(LDFLAGS)" github.com/spatialcurrent/gotmpl/cmd/gotmpl

bin/gotmpl_windows_amd64.exe:
	GOOS=windows GOARCH=amd64 go build -o bin/gotmpl_windows_amd64.exe -gcflags="$(GCFLAGS)" -ldflags="$(LDFLAGS)" github.com/spatialcurrent/gotmpl/cmd/gotmpl

bin/gotmpl_linux_arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/gotmpl_linux_arm64 -gcflags="$(GCFLAGS)" -ldflags="$(LDFLAGS)" github.com/spatialcurrent/gotmpl/cmd/gotmpl

build_cli: bin/gotmpl_darwin_amd64 bin/gotmpl_linux_amd64 bin/gotmpl_windows_amd64.exe bin/gotmpl_linux_arm64  ## Build command line programs

#
# Shared Objects
#

bin/gotmpl.so:  ## Compile Shared Object for current platform
	# https://golang.org/cmd/link/
	# CGO Enabled : https://github.com/golang/go/issues/24068
	CGO_ENABLED=1 go build -o $(DEST)/gotmpl.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/gotmpl/plugins/gotmpl

bin/gotmpl_linux_amd64.so:  ## Compile Shared Object for Linux / amd64
	# https://golang.org/cmd/link/
	# CGO Enabled : https://github.com/golang/go/issues/24068
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(DEST)/gotmpl_linux_amd64.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/gotmpl/plugins/gotmpl

bin/gotmpl_linux_armv7.so:  ## Compile Shared Object for Linux / ARMv7
	# LDFLAGS - https://golang.org/cmd/link/
	# CGO Enabled  - https://github.com/golang/go/issues/24068
	# GOARM/GOARCH Compatability Table - https://github.com/golang/go/wiki/GoArm
	# ARM Cross Compiler Required - https://www.acmesystems.it/arm9_toolchain
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -ldflags "-linkmode external -extldflags -static" -o $(DEST)/gotmpl_linux_armv7.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/gotmpl/plugins/gotmpl

bin/gotmpl_linux_armv8.so:   ## Compile Shared Object for Linux / ARMv8
	# LDFLAGS - https://golang.org/cmd/link/
	# CGO Enabled  - https://github.com/golang/go/issues/24068
	# GOARM/GOARCH Compatability Table - https://github.com/golang/go/wiki/GoArm
	# ARM Cross Compiler Required - https://www.acmesystems.it/arm9_toolchain
	# Dependencies - https://www.96boards.org/blog/cross-compile-files-x86-linux-to-96boards/
	GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc go build -ldflags "-linkmode external -extldflags -static" -o $(DEST)/gotmpl_linux_armv8.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/gotmpl/plugins/gotmpl

build_so: bin/gotmpl_linux_amd64.so bin/gotmpl_linux_armv7.so bin/gotmpl_linux_armv8.so  ## Build Shared Objects (.so)

#
# JavaScript
#

dist/gotmpl.mod.js:  ## Build JavaScript module
	gopherjs build -o dist/gotmpl.mod.js github.com/spatialcurrent/gotmpl/cmd/gotmpl.mod.js

dist/gotmpl.mod.min.js:  ## Build minified JavaScript module
	gopherjs build -m -o dist/gotmpl.mod.min.js github.com/spatialcurrent/gotmpl/cmd/gotmpl.mod.js

dist/gotmpl.global.js:  ## Build JavaScript library that attaches to global or window.
	gopherjs build -o dist/gotmpl.global.js github.com/spatialcurrent/gotmpl/cmd/gotmpl.global.js

dist/gotmpl.global.min.js:  ## Build minified JavaScript library that attaches to global or window.
	gopherjs build -m -o dist/gotmpl.global.min.js github.com/spatialcurrent/gotmpl/cmd/gotmpl.global.js

build_javascript: dist/gotmpl.mod.js dist/gotmpl.mod.min.js dist/gotmpl.global.js dist/gotmpl.global.min.js  ## Build artifacts for JavaScript

test_javascript:  ## Run JavaScript tests
	npm run test

lint:  ## Lint JavaScript source code
	npm run lint

#
# Examples
#

bin/gotmpl_example_c: bin/gotmpl.so  ## Build C example
	mkdir -p bin && cd bin && gcc -o gotmpl_example_c -I. ./../examples/c/main.c -L. -l:gotmpl.so

bin/gotmpl_example_cpp: bin/gotmpl.so  ## Build C++ example
	mkdir -p bin && cd bin && g++ -o gotmpl_example_cpp -I . ./../examples/cpp/main.cpp -L. -l:gotmpl.so

run_example_c: bin/gotmpl.so bin/gotmpl_example_c  ## Run C example
	cd bin && LD_LIBRARY_PATH=. ./gotmpl_example_c

run_example_cpp: bin/gotmpl.so bin/gotmpl_example_cpp  ## Run C++ example
	cd bin && LD_LIBRARY_PATH=. ./gotmpl_example_cpp

run_example_python: bin/gotmpl.so  ## Run Python example
	LD_LIBRARY_PATH=bin python examples/python/test.py

run_example_javascript: dist/gotmpl.mod.js  ## Run JavaScript module example
	node examples/js/index.mod.js

#
# Clean
#

clean:
	rm -fr bin
	rm -fr dist
