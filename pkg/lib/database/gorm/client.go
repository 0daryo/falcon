package gorm

import (
	"github.com/0daryo/falcon/pkg/lib/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver for gorm
)

// NewSQLClient ... get mysql client
func NewSQLClient(cfg *database.Config) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", cfg.OutputDBConfig().FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
