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
