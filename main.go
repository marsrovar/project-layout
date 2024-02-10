package main

import (
	"iserver/initialize"
	"iserver/server"
	"iserver/utils/env"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	env.Env = env.Get("Env")

	initialize.InitStorage()

	server.StartServer(env.Get("Port"))
}
