package helpers

import (
	"github.com/MantisSTS/BountyProcess/server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DatabaseHelper struct {
	Database *gorm.DB
}

func (db *DatabaseHelper) Connect() *gorm.DB {

	var err error
	db.Database, err = gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.Migrate()

	return db.Database
}

func (db *DatabaseHelper) Migrate() {
	db.Database.AutoMigrate(&models.Domain{})
	db.Database.AutoMigrate(&models.Subdomain{})
	// db.Database.AutoMigrate(&models.WAF{})
	// db.Database.AutoMigrate(&models.Crawler{})
	// db.Database.AutoMigrate(&models.Screenshot{})
	// db.Database.AutoMigrate(&models.Whois{})
	// db.Database.AutoMigrate(&models.DNS{})
	// db.Database.AutoMigrate(&models.Nmap{})
	// db.Database.AutoMigrate(&models.SSL{})
	// db.Database.AutoMigrate(&models.HTTP{})
	// db.Database.AutoMigrate(&models.JS{})
	// db.Database.AutoMigrate(&models.Nuclei{})
	// db.Database.AutoMigrate(&models.Wayback{})
	// db.Database.AutoMigrate(&models.Github{})
}
