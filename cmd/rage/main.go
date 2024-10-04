package main

import (
	"flag"
	"log"

	"github.com/klystro/rage/pkg/utils"
)

func main() {
	var entryPoint string
	flag.StringVar(&entryPoint, "e", "", "entry point to run (e.g., cmd/klystro/main.go)")
	flag.Parse()

	// Default configuration file
	configFile := "rage.yaml"
	if entryPoint == "" {
		entryPoint = utils.GetEntryPointFromConfig(configFile)
		if entryPoint == "" {
			log.Fatalf("No entry point specified and no rage.yaml found")
		}
	}
}
