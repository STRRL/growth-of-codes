package api

import (
	"net/http"

	"github.com/STRRL/growth-of-codes/pkg/api/serverless"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
	app = serverless.NewServerlessGinApp()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
