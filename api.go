package main

import (
	"net/http"
	"path"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const V1 = "v1"
const BOOK = "book"

type Meta struct {
	Authors []struct {
		Name string `json:"name"`
	} `json:"authors"`
	Tags []struct {
		Name string `json:"name"`
	} `json:"tags"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Date     string `json:"date"`
	Language string `json:"language"`
	Rights   string `json:"rights"`
}

type Interior struct {
	Payload       string   `json:"payload"`
	SectionTitles []string `json:"sectionTitles"`
}

type EBookMiddle struct {
	Meta     `json:"meta"`
	Interior `json:"interior"`
}

func initAPI() {
	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	r.Use(cors.Default())

	r.POST(path.Join(V1, BOOK), func(c *gin.Context) {
		var ebook EBookMiddle

		if err := c.BindJSON(&ebook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		err := createHTMLBook(ebook)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to create book",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	r.Run()

}
