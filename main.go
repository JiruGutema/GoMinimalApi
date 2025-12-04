package main

import (
	"fmt"

	"github.com/jirugutema/gominimalapi/config"
	"github.com/jirugutema/gominimalapi/routes"
)

func main() {
	// run the server
	config.ConnectDatabase()
	routes.Router()
	string := fmt.Sprintf("Server is running on port: %s!", "8080")
	fmt.Println(string)
}
