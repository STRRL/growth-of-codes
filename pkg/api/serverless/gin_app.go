package serverless

import "github.com/gin-gonic/gin"

func NewServerlessGinApp() *gin.Engine {
	app := gin.New()
	router := app.Group("/api")
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	return app
}
