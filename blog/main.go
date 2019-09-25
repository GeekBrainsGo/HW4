package main

import (
	"blog/app/webserver"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/blog.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := webserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := webserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
