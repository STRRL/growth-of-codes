package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()
	r := app.Group("/api")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
