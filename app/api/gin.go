package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var app *gin.Engine

func init() {
	app = gin.New()
	router := app.Group("/api")
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)

}
