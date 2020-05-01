package main

import (
	"log"

	"github.com/mbertschler/serverman/pkg"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Serverman started")
	log.Println(pkg.VersionString())
}
