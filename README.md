
Markdown memo tool


# Build

    go build -o bin/gopad src/*.go

move directry "bin"


# Run

    gopad -p port filename

default port 5005
default filename gopad.db

# Run(Develop)

    go run src/*.go -base bin

default base ""

