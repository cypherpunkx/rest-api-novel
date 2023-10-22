package main

import (
	"enigmacamp.com/rest-api-novel/config"
	"enigmacamp.com/rest-api-novel/delivery"
)

func init() {
	config.InitiliazeConfig()
	config.InitDB()
}

func main() {
	delivery.Server().Run()
}
