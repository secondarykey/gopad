
Markdown memo tool

# Using





# Build

    go build -o bin/gopad src/*.go

move directry "bin"


# Run

    gopad -p port filename

default port 5005
default filename gopad.db

# If you publish if


    gopad -p port -server "" filename

or

    gopad -p port -server "" filename

# Run(Develop)

    go run src/*.go -base bin

default base ""

