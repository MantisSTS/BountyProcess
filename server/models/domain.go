package models

import (
	"github.com/jinzhu/gorm"
)

// Domain model
type Domain struct {
	gorm.Model
	Domain string `json:"domain"`
	// Subdomains []Subdomain `json:"subdomains"`
	// WAFs       []WAF       `json:"wafs"`
	// Crawlers   []Crawler   `json:"crawlers"`
	// Screenshots []Screenshot `json:"screenshots"`
	// Whois      []Whois      `json:"whois"`
	// DNS        []DNS        `json:"dns"`
	// Nmap       []Nmap       `json:"nmap"`
	// SSL        []SSL        `json:"ssl"`
	// HTTP       []HTTP       `json:"http"`
	// JS         []JS         `json:"js"`
	// Nuclei     []Nuclei     `json:"nuclei"`
	// Wayback    []Wayback    `json:"wayback"`
	// Github     []Github     `json:"github"`
}
