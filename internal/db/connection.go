package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dangLuan01/karaoke/internal/config"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *goqu.Database

func InitDB() error {

	connStr := config.NewConfig().DNS()
	var err error
    //DB, err = sql.Open("mysql", connStr)
	sqlDB, err := sql.Open("mysql", connStr)
    if err != nil {
        log.Fatal(err)
    }
	//defer sqlDB.Close()

	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil{
		sqlDB.Close()
		return fmt.Errorf("DB ping error: %s", err)
	}
	DB = goqu.New("mysql", sqlDB)

	log.Println("Connected")

	return nil
}