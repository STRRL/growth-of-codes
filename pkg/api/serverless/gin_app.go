package serverless

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewServerlessGinApp() *gin.Engine {
	app := gin.New()
	router := app.Group("/api")
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.GET("/random/:count", func(c *gin.Context) {
		count := defaultCount
		countParameter := c.Param("count")
		count, _ = strconv.Atoi(countParameter)
		c.JSON(200, RandomTimeSeries(count))
	})
	router.GET("/repo/chaos-mesh", func(c *gin.Context) {
		result, err := ComplexityOfChaosMesh()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, result)
	})
	return app
}
