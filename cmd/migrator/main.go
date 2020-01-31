package main

import "Multi-Honeypot/internal/app/backend"

func main() {
	app := backend.NewApp("/Users/hackboy/Code/Multi-Honeypot/config/config.ini")
	err := app.DB.Migrator()
	if err != nil {
		panic(err)
	}
}
