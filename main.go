package main

import (
	"fmt"
	"gominimalapi/routes"

)


func main() {
	//run the server
	routes.Router()
	string := fmt.Sprintf("Server is running on port: %s!", "8080")
	fmt.Println(string)
}
