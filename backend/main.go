package main

import (
	"VRisingAcademySite/api"
	"VRisingAcademySite/database"
	"fmt"
	"os"
	"strconv"

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
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"https://vrising-academy.info"}

	r.Use(cors.Default())

	api.RegisterApiHandlers(r)

	//TLS
	// crtPath := "ssl/"
	// r.RunTLS(":10443", crtPath+"combined.crt", crtPath+"private.key")

	//HTTP
	var port int = 8087
	if database.IsTestMode() {
		port = 4047
	}
	r.Run("localhost:" + strconv.Itoa(port))
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--test-mode" {
			fmt.Println("Launching Academy server in TEST mode...")
			for i := 0; i < 30; i++ {
				fmt.Println("\x1B[31m!!!!! YOU RUN ACADEMY SERVER IN TEST MODE, PLEASE ENSURE THIS IS WHAT YOU WANTED !!!!!")
			}
			fmt.Println("\x1B[0m")
			database.SetTestMode()
		}
	}
	handleRequest()
}
