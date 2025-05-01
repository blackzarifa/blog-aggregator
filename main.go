package main

import (
	"fmt"
	"gator/internal/config"
)

const configFileName = ".gatorconfig.json"

func main() {
	cfg, err := config.Read(configFileName)
	if err != nil {
		panic(err)
	}

	cfg.SetUser("lane")

	cfg, err = config.Read(configFileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
