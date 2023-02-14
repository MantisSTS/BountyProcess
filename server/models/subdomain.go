package models

import (
	"github.com/jinzhu/gorm"
)

// Domain model
type Subdomain struct {
	gorm.Model
	ParentID  uint   `json:"domain_id"`
	Subdomain string `json:"subdomain"`
}
