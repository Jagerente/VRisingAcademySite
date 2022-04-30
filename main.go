package main

import (
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
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("resources/", "./resources")

	r.GET("/", home_page)

	//TLS
	// crtPath := "ssl/"
	// r.RunTLS(":10443", crtPath+"combined.crt", crtPath+"private.key")

	//HTTP
	r.Run(":8080")
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
