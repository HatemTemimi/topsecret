package main

import (
	"server/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {

	echo := echo.New()
	server := &server.Server{}
	server.SetupAndLaunch(echo)

}
