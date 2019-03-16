package main

import (
	"flag"
	"github.com/kangaechu/godcast/pkg/godcast"
)

func main() {
	confFile := flag.String("c", "podcast.yaml", "conf yaml file name")
	flag.Parse()
	godcast.Run(*confFile)
}
