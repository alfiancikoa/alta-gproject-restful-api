package main

import (
	"alte/e-commerce/config"
	"alte/e-commerce/routers"
)

func main() {
	// Configuration to Database
	config.InitDB()
	// Call the router
	e := routers.New()
	// Logger to run server with port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
