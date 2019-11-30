package main

import (
	"flag"
	"fmt"
	"github.com/kangaechu/godcast"
	"os"
)

var version string
var revision string

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf(os.Args[0]+": %s-%s\n", version, revision)
		os.Exit(0)
	}

	confFile := flag.String("c", "podcast.yaml", "conf yaml file name")
	flag.Parse()
	godcast.Run(*confFile)
}
