package main

import (
	"fmt"

	"api-merchant-backend/routers/server"
)

func main() {
	app := server.Server()

	err := app.Start(":8080")
	if err != nil {
		fmt.Println("Error with : ", err)
	}
}
