package main

import (
	"server/internal/server"
)

func main() {
	server := server.Server{}
	server.SetupAndLaunch()
}
