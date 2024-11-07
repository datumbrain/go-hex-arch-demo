package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Client struct {
	db *gorm.DB
}

func New(url string) *Client {
	var err error
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	return &Client{db: db}
}

func (c *Client) DB() *gorm.DB {
	return c.db
}

func (c *Client) Close() error {
	sql, err := c.db.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}
