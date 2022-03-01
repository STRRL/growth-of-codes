package main

import "github.com/STRRL/growth-of-codes/pkg/api/serverless"

func main() {
	app := serverless.NewServerlessGinApp()
	app.Run()
}
