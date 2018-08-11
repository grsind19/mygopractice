package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("Parsing failed")
	}
	fmt.Println(config.Get("path"))
}
