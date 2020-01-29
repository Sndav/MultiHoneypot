package main

import "Muti-Honeypot/internal/app/backend"

func main(){
	app := backend.NewApp("/Users/hackboy/Code/Muti-Honeypot/config/config.ini")
	app.Start()
}


