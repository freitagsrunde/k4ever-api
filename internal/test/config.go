package test

import (
	"fmt"
	"os"

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

func (c *Config) Version() string {
	return "1.0"
}

func (c *Config) GitCommit() string {
	return "commit"
}

func (c *Config) GitBranch() string {
	return "test"
}

func (c *Config) BuildTime() string {
	return "now"
}

func (c *Config) FilesPath() string {
	return "/temp/k4ever/"
}

func (c *Config) DB() *gorm.DB {
	if c.db == nil {
		c.connectToDatabase()
	}
	return c.db
}

func (c *Config) LdapHost() string {
	return "localhost"
}

func (c *Config) LdapBind() string {
	return "admin"
}

func (c *Config) LdapPassword() string {
	return "admin"
}

func (c *Config) LdapBaseDN() string {
	return "CN=Users,DC=example,DC=com"
}

func (c *Config) LdapFilterDN() string {
	return "(&(objectClass=person)(uid={username}))"
}

func (c *Config) HttpServerHost() string {
	return "localhost"
}

func (c *Config) SetHttpServerPort(port int) {
	return
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
		&models.History{},
		&models.PurchaseItem{},
	)

	deposit := models.Product{}
	deposit.Name = "Deposit"
	deposit.Hidden = true
	if err := db.Where(deposit).FirstOrCreate(&deposit).Error; err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
