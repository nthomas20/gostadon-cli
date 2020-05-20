# This how we want to name the binary output
BINARY=gostadon-cli

# These are the values we want to pass for VERSION and BUILD ( Semantic Versioning Recommended: https://semver.org/ )
VERSION=`git describe --tags --abbrev=0`
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.version=${VERSION} -X main.buildDate=${BUILD}"

# Builds the project ( https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04 )
build:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY}
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY}.exe
	env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY}-darwin

# Installs our project: copies binaries
install:
	go install ${LDFLAGS}

# Cleans our project: deletes binaries
clean:
	if [ -f bin/${BINARY} ] ; then rm bin/${BINARY} ; fi
	if [ -f bin/${BINARY}.exe ] ; then rm bin/${BINARY}.exe ; fi
	if [ -f bin/${BINARY}-darwin ] ; then rm bin/${BINARY}-darwin ; fi

# Pushes release of current tag including build artifact ( Requires https://github.com/github/hub )
release:
	git-chglog -o CHANGELOG.md
	git commit -m "changelog" CHANGELOG.md
	git push
	git push origin ${VERSION}
	hub release create -a "bin/${BINARY}#gostadon-cli (Linux-amd64)" -a "bin/${BINARY}.exe#gostadon-cli.exe (Windows-amd64)" -a "bin/${BINARY}-darwin#gostadon-cli-darwin (MacOS-amd64)" -m "gostadon-cli ${VERSION}" -m "Go Client for encrypted messaging" ${VERSION}

.PHONY: clean install