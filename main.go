package main

import (
	"VRisingAcademySite/api"
	"VRisingAcademySite/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Page struct {
// 	Title string `json:"title"`
// 	Url   string `json:"url"`
// }

func home_page(c *gin.Context) {
	// pages := initPages()
	c.HTML(http.StatusOK, "home_page.html", nil)
}

func handleRequest() {
	if database.CheckIfDatabaseExists() == false {
		database.InitializeDatabase()
		fmt.Println("No initialized database found. Initializing new...")
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("resources/", "./resources")

	api.RegisterApiHandlers(r)
	r.GET("/", home_page)

	//TLS
	// crtPath := "ssl/"
	// r.RunTLS(":10443", crtPath+"combined.crt", crtPath+"private.key")

	//HTTP
	r.Run("localhost:8080")
}

func main() {
	handleRequest()
}

// func initPages() []Page {
// 	content, err := ioutil.ReadFile("./pages.json")
// 	if err != nil {
// 		log.Fatal("Error when opening file: ", err)
// 	}

// 	var payload []Page
// 	err = json.Unmarshal(content, &payload)
// 	if err != nil {
// 		log.Fatal("Error during Unmarshal(): ", err)
// 	}
// 	return payload
// }
