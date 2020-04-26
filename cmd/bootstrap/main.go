package main

import (
	"log"

	"github.com/mbertschler/serverman/pkg"
	"github.com/mbertschler/serverman/pkg/golang"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("hello world")

	log.Println(pkg.VersionString())

	err := golang.Install()
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
