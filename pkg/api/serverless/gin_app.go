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
	router.GET("/complexity", func(c *gin.Context) {
		repo := c.Query("repo")
		language := c.Query("language")
		if len(language) == 0 || len(repo) == 0 {
			c.JSON(200, nil)
			return
		}
		result, err := ComplexityForRepositoryAndLanguage(repo, language)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, result)
	})
	router.GET("/repo/list", func(c *gin.Context) {
		result, err := AllAvailableRepo()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, result)
	})
	router.GET("/language/list", func(c *gin.Context) {
		repo := c.Query("repo")
		result, err := AvailableLanguageForRepo(repo)
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
