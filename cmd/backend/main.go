package main

import "Multi-Honeypot/internal/app/backend"

func main() {
	app := backend.NewApp("/Users/hackboy/Code/Multi-Honeypot/config/config.ini")
	app.Start()
}
