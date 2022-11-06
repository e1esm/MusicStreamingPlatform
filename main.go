package main

import (
	"musicStreamingPlatform/config"
)

func main() {
	config := config.GetConfig()

	app := &App{}
	app.Initialize(config)
	app.Run(":823")
}
