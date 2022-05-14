package main

import (
	"project/config"
	"project/routes"
)

func main() {
	config.Init()
	e := routes.New()
	e.Start(":8802")
}
