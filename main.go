package main

import (
	"VRisingAcademySite/api"
	"VRisingAcademySite/database"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func handleRequest() {
	if !database.CheckIfDatabaseExists() {
		database.InitializeDatabase()
		fmt.Println("No initialized database found. Initializing new...")
	}
	if database.CheckIfDatabaseNeedsUpdate() {
		fmt.Println(fmt.Sprintf(`Obsolete database version found. Running update to version %d`, database.DatabaseVersion))
		database.UpdateDatabase()
	}

	r := gin.Default()

	// CORS FULL CONFIG EXAMPLE
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://foo.com"},
	// 	AllowMethods:     []string{"PUT", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 	  return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	//   }))

	//CORS DEFAULT
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}

	r.Use(cors.New(config))

	api.RegisterApiHandlers(r)

	//TLS
	// crtPath := "ssl/"
	// r.RunTLS(":10443", crtPath+"combined.crt", crtPath+"private.key")

	//HTTP
	r.Run("localhost:8081")
}

func main() {
	handleRequest()
}
