package main

import "avitotest/internal/app"

const configPath = "./config/.env"

func main() {
	app.Run(configPath)
}
