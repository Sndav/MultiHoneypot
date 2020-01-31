package main

import (
	"Multi-Honeypot/internal/app/website"
)

func main() {
	app := website.NewServer("/Users/hackboy/Code/Multi-Honeypot/config/config.ini")
	app.Start()
}
