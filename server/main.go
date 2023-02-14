package main

import (
	"net/http"

	controllers "github.com/MantisSTS/BountyProcess/server/controllers"
	helpers "github.com/MantisSTS/BountyProcess/server/helpers"
	"github.com/gin-gonic/gin"
)

func main() {

	// Create all the necessary queues
	rmq := helpers.RabbitMQHelper{}
	rmq.CreateQueue("init", "domain")
	rmq.CreateQueue("init", "subdomain")
	rmq.CreateQueue("init", "waf")
	rmq.CreateQueue("init", "crawler")
	rmq.CreateQueue("init", "screenshot")
	rmq.CreateQueue("init", "whois")
	rmq.CreateQueue("init", "dns")
	rmq.CreateQueue("init", "nmap")
	rmq.CreateQueue("init", "ssl")
	rmq.CreateQueue("init", "http")
	rmq.CreateQueue("init", "js")
	rmq.CreateQueue("init", "nuclei")
	rmq.CreateQueue("init", "wayback")
	rmq.CreateQueue("init", "github")

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
			domain.PUT("/subdomain/create", dc.CreateSubdomain)
		}
	}

	r.Run(":6123")
}
