package main

import (
	"net/http"

	controllers "github.com/MantisSTS/BountyProcess/server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// Add a global middleware to the router
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		// Ping test
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		domain := v1.Group("/domain")
		{
			dc := controllers.DomainController{}
			domain.PUT("/create", dc.Create)
		}
	}

	r.Run(":6123")
}
