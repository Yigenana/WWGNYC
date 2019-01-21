# Hands On Demo Instructions

## Download Docker

This demo will be run in Docker so we can run Go 1.11.

[Install on Mac](https://docs.docker.com/docker-for-mac/install/)
[Install on Windows](https://docs.docker.com/docker-for-windows/install/)

## Run Docker with Go

`docker --rm --it golang:1.11`

## Creating a Go program

Move out of go path:

`cd`

`go env GOPATH`

Create a new directory and navigate into it:

`mkdir mycat`

`cd mycat`

Install vim in order to edit files:

`apt-get update && apt-get install -y vim`

Create main Go file:

`vim main.go`

Edit file:

```go
package main

import (
    "io"
    "os"

    "github.com/sirupsen/logrus"
)

func main() {
    _, err := io.Copy(os.Stdout, os.Stdin)
    if err != nil {
        logrus.Fatal(err)
    }
}
```

## Setting Up Go Modules

Run the file:

`go run main.go`

Notice the error.

Let's try to get modules:

`go get`

Go can't find the modules becase we're not in Go Path.
Go modules to the rescue!

Enable go modules:

`go mod init`

Again!

`go mod init github.com/yigenana/mycat`

## Working with Go Modules

A go.mod file is created:

`cat go.mod`

Build the module. Go modules will now fetch our dependencies.

`go build`

Confirm that the file runs.

`./mycat`

`hello`

Check go.mod again, see the dependecies. Also notices the new file go.sum.

`cat go.mod`

`cat go.sum`

## Changing Dependencies

Let's change the version of a dependency:

`vi go.mod`

`github.com/sirupsen/logrus v1.2.0`

`go build`

Review the go.sum file, notice both versions are still in the file.

`cat go.sum`

Run the tidy command

`go mod tidy`

Tidy adds the requirements for the test files, which are ignored with build. This ensures your tests will pass.

Understand your dependencies:

`go mod why -m <module>`

## Backwards compatability

Go modules supports the vendor directory so that your code can be supported by Go programs no yet using modules.

`go mod vendor`
