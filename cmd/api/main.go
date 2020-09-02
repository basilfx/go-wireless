package main

import (
	"github.com/basilfx/go-wireless/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.SetupRoutes(r)
	r.Run(":8080")
}
