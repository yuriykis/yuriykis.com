package main

import (
	"log"
	"os"

	"github.com/yuriykis/yuriykis.com/cmd/server"
)

func main() {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	server.Run(rootPath)
}
