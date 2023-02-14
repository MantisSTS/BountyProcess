package controllers

import (
	"fmt"
	"net/http"
	"sync"

	helpers "github.com/MantisSTS/BountyProcess/server/helpers"
	"github.com/gin-gonic/gin"
)

type DomainController struct{}

func (dc *DomainController) Create(c *gin.Context) {
	var resp HttpStdResponse

	// get the request body
	var req struct {
		Domains []string `json:"domains"`
	}

	if err := c.BindJSON(&req); err != nil {

		resp.Success = false
		resp.Data = "Invalid request"

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var rmq helpers.RabbitMQHelper
	var wg sync.WaitGroup

	for i, domain := range req.Domains {
		go func(counter int, domain string) {
			wg.Add(1)
			err := rmq.Publish(fmt.Sprintf("chan.%d.%s", counter, domain), "domain", domain)

			if err != nil {
				resp.Success = false
				resp.Data = "Invalid Request"
				c.JSON(http.StatusInternalServerError, resp)
			}
		}(i, domain)
	}
	wg.Wait()

	resp.Success = true
	resp.Data = req

	c.JSON(http.StatusOK, resp)
}

func (dc *DomainController) CreateSubdomain(c *gin.Context) {
	var resp HttpStdResponse

	// get the request body
	var req struct {
		Domains []string `json:"domains"`
	}

	if err := c.BindJSON(&req); err != nil {

		resp.Success = false
		resp.Data = "Invalid request"

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	var rmq helpers.RabbitMQHelper
	var wg sync.WaitGroup

	for i, domain := range req.Domains {
		go func(counter int, domain string) {
			wg.Add(1)
			// Add to a database
		}(i, domain)
	}
	wg.Wait()

	resp.Success = true
	resp.Data = req

	c.JSON(http.StatusOK, resp)
}
