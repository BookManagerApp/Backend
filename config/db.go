package config

import (
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func LoadEnv() {
    err := godotenv.Load("/home/dwidt/myapp/.env")
    if err != nil {
        panic("Error loading .env file")
    }
}

func CreateDBConnection() *gorm.DB {
    LoadEnv()

    // string connection database
    dbConfig := os.Getenv("SQLSTRING")

    // koneksi ke database
    DB, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
    })
    if err != nil {
        panic(err)
    }

    dbO, err := DB.DB()
    if err != nil {
        panic(err)
    }
    dbO.SetConnMaxIdleTime(time.Duration(1) * time.Minute)
    dbO.SetMaxIdleConns(2)
    dbO.SetConnMaxLifetime(time.Duration(1) * time.Minute)

    DB.Statement.RaiseErrorOnNotFound = true

    return DB
}
