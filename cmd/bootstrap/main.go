package main

import (
	"log"

	"github.com/mbertschler/serverman/pkg/version"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("hello world")

	log.Println(version.String())
}
