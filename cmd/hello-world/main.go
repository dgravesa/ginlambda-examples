package main

import (
	"net/http"

	"github.com/dgravesa/ginlambda"
	"github.com/gin-gonic/gin"
)

func helloHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello, World!")
}

func main() {
	r := gin.Default()
	r.GET("/greeting", helloHandler)

	ginlambda.Start(r)
}
