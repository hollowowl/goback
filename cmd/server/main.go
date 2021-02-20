package main

import (
	"arduino-playground.xyz/goback/config"
	"arduino-playground.xyz/goback/internal"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("Usage: go run cmd/server/main.go config_file.json\nExample: go run cmd/server/main.go config/config-test.json")
	}
	configFile := args[0]
	config, err := config.FromJson(configFile)
	if err != nil {
		log.Fatalln("Config parse failed: ", err.Error(), configFile)
	}
	server, err := internal.NewServer(config)
	if err != nil {
		log.Fatalln("Server init failed", err.Error())
	}
	log.Println("Staring server...")
	server.Run()
}
