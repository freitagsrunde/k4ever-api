package context

import (
	"github.com/freitagsrunde/k4ever-backend/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

type Config struct {
	appVersion     string
	db             *gorm.DB
	httpServerPort int
}

func NewConfig() *Config {
	c := &Config{}
	c.appVersion = viper.GetString("version")
	c.httpServerPort = viper.GetInt("port")

	return c
}

func (c *Config) AppVersion() string {
	return c.appVersion
}

func (c *Config) DB() *gorm.DB {
	if c.db == nil {
		c.connectToDatabase()
	}
	return c.db
}

func (c *Config) HttpServerPort() int {
	return c.httpServerPort
}

func (c *Config) connectToDatabase() error {
	db, err := gorm.Open("sqlite3", "test.db")
	c.db = db

	return err
}

func (c *Config) MigrateDB() {
	db := c.DB()

	db.AutoMigrate(
		&models.Product{},
	)
}
