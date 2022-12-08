package main

import (
	"AdvancedNetwork/connections"
	"AdvancedNetwork/pkg/apis"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Starting UI Generator Application")
	connections.ConnectMongo()

	e := echo.New()
	api.EchoManager(e)
	e.Logger.Fatal(e.Start(":8088"))
}
