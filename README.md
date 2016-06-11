
Markdown memo tool

# Using

    github.com/mattn/go-sqlite3
    github.com/monochromegane/argen

# Build

    go build -o bin/gopad src/*.go

# Run

move directry "bin"

    gopad -p port filename

default port 5005
default filename gopad.db

# If you publish if

    gopad -server "" filename

default localhost

# Run(Develop)

    go run src/*.go -base bin

default base ""

