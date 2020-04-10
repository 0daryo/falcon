package database

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

var timeZoneJST = time.FixedZone("Asia/Tokyo", 9*60*60)

type Config struct {
	DBHost     string
	DBPort     int64
	DBNet      string
	DBUser     string
	DBPassword string
	DBName     string
}

func (c *Config) OutputDBConfig() *mysql.Config {
	return &mysql.Config{
		User:                 c.DBUser,
		Passwd:               c.DBPassword,
		Addr:                 c.DBHost,
		Net:                  c.DBNet,
		DBName:               c.DBName,
		Loc:                  timeZoneJST,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
}
