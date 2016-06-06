
Markdown memo tool


# Build

    go build -o bin/gopad.exe src/*.go

    cd bin

    gopad.exe -p port "database filename(default gopad.db)"

default port 5005

# Run(Develop)

    go run src/*.go -base bin
