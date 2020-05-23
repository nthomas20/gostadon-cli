# This how we want to name the binary output
BINARY_LINUX=gostadon-cli
BINARY_SNAP=gostadon-cli
BINARY_WINDOWS=gostadon-cli.exe
BINARY_MACOS=gostadon-darwin
DESCRIPTOR_LINUX=Linux-amd64
DESCRIPTOR_WINDOWS=Windows-amd64
DESCRIPTOR_MACOS=MacOS

RELEASE_TITLE=gostadon-cli
RELEASE_MESSAGE=Go Client for encrypted messaging

# These are the values we want to pass for VERSION and BUILD ( Semantic Versioning Recommended: https://semver.org/ )
VERSION=`git describe --tags --abbrev=0`
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.version=${VERSION} -X main.buildDate=${BUILD}"
LDFLAGS_SNAP=-ldflags "-w -s -X main.version=${VERSION}-snap -X main.buildDate=${BUILD}"
LDFLAGS_HERE=-ldflags "-w -s -X main.version=${VERSION}-local -X main.buildDate=${BUILD}"

# Builds the project ( https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04 )
build:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_LINUX}
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_WINDOWS}
	env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_MACOS}

build-snap:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS_SNAP} -o bin/${BINARY_SNAP}

# Installs our project: copies binaries
install:
	# Use the "here" build version to presentationally disconnect this from an official build release
	go install ${LDFLAGS_HERE}

# Cleans our project: deletes binaries
clean:
	if [ -f bin/${BINARY_LINUX} ] ; then rm bin/${BINARY_LINUX} ; fi
	if [ -f bin/${BINARY_SNAP} ] ; then rm bin/${BINARY_SNAP} ; fi
	if [ -f bin/${BINARY_WINDOWS} ] ; then rm bin/${BINARY_WINDOWS} ; fi
	if [ -f bin/${BINARY_MACOS} ] ; then rm bin/${BINARY_MACOS} ; fi

# Generate and push changelog
changelog:
	git pull
	git-chglog -o CHANGELOG.md
	git commit -a -m "changelog"
	git push
	git push origin ${VERSION}

# Pushes release of current tag including build artifact ( Requires https://github.com/github/hub )
release:
	hub release create \
		-a "bin/${BINARY_LINUX}#${BINARY_LINUX} (${DESCRIPTOR_LINUX})" \
		-a "bin/${BINARY_WINDOWS}#${BINARY_WINDOWS} (${DESCRIPTOR_WINDOWS})" \
		-a "bin/${BINARY_MACOS}#${BINARY_MACOS} (${DESCRIPTOR_MACOS})" \
		-m "${RELEASE_TITLE} ${VERSION}" \
		-m "${RELEASE_MESSAGE}" \
		${VERSION}

.PHONY: clean build