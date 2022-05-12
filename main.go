package main

import (
	"VRisingAcademySite/api"
	"VRisingAcademySite/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func handleRequest() {
	if !database.CheckIfDatabaseExists() {
		database.InitializeDatabase()
		fmt.Println("No initialized database found. Initializing new...")
	}

	r := gin.Default()

	api.RegisterApiHandlers(r)

	//TLS
	// crtPath := "ssl/"
	// r.RunTLS(":10443", crtPath+"combined.crt", crtPath+"private.key")

	//HTTP
}

func main() {
	handleRequest()
}
