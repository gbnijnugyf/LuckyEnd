package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shawu21/test/router"
)

func main() {
	r := gin.Default()
	router.Routers(r)
	r.Run()
}
