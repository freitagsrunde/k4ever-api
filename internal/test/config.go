package test

import (
	"github.com/freitagsrunde/k4ever-backend/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	conf := NewConfig()
	conf.MigrateDB()
}

type Config struct {
	db *gorm.DB
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) AppVersion() string {
	return "1.0"
}

func (c *Config) DB() *gorm.DB {
	if c.db == nil {
		c.connectToDatabase()
	}
	return c.db
}

func (c *Config) HttpServerPort() int {
	return 8080
}

func (c *Config) connectToDatabase() error {
	db, err := gorm.Open("sqlite3", ":memory:")
	c.db = db

	return err
}

func (c *Config) MigrateDB() {
	db := c.DB()

	db.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.Permission{},
		&models.Purchase{},
		&models.PurchaseItem{},
	)
}
