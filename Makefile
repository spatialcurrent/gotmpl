# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

bin/gotmpl_darwin_amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/gotmpl_darwin_amd64 $$(go list ./...)

bin/gotmpl_linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/gotmpl_linux_amd64 $$(go list ./...)

bin/gotmpl_windows_amd64.exe:
	GOOS=windows GOARCH=amd64 go build -o bin/gotmpl_windows_amd64.exe $$(go list ./...)

bin/gotmpl_linux_arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/gotmpl_linux_arm64 $$(go list ./...)

build: \
bin/gotmpl_darwin_amd64 \
bin/gotmpl_linux_amd64 \
bin/gotmpl_windows_amd64.exe \
bin/gotmpl_linux_arm64

fmt:
	go fmt $$(go list ./... )

install:
	go install $$(go list ./...)

test:
	bash scripts/test.sh

clean:
	rm -fr bin
