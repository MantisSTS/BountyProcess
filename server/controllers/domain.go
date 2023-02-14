package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	helpers "github.com/MantisSTS/BountyProcess/server/helpers"
	"github.com/MantisSTS/BountyProcess/server/models"
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

	results := make(chan string, len(req.Domains))

	for i, domain := range req.Domains {
		go func(counter int, domain string) {
			wg.Add(1)
			err := rmq.Publish(fmt.Sprintf("chan.%d.%s", counter, domain), "domain", domain)
			if err != nil {
				resp.Success = false
				resp.Data = "Invalid Request"
				c.JSON(http.StatusInternalServerError, resp)
			}
			results <- strings.TrimSpace(domain)
		}(i, domain)
	}

	wg.Wait()

	go func() {
		for res := range results {
			// Add to the database
			x := helpers.DatabaseHelper{}
			db := x.Connect()
			defer db.Close()

			var domainModel models.Domain
			domainModel.Domain = strings.TrimSpace(res)

			db.Create(domainModel)
		}
	}()

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

	results := make(chan string, len(req.Domains))

	for _, domain := range req.Domains {
		results <- strings.TrimSpace(domain)
		// Add to a database
	}

	go func() {
		for res := range results {
			// Add to the database
			x := helpers.DatabaseHelper{}
			db := x.Connect()
			defer db.Close()

			var d models.Domain
			db.Where("domain = ?", res).First(&d)
			parentId := d.ID

			var subdomainModel models.Subdomain
			subdomainModel.ParentID = parentId
			subdomainModel.Subdomain = strings.TrimSpace(res)

			db.Create(&subdomainModel)
		}
	}()

	resp.Success = true
	resp.Data = req

	c.JSON(http.StatusOK, resp)
}
