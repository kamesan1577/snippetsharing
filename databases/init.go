package databases

import (
	"database/sql"
	"github.com/joho/godotenv"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	err = godotenv.Load()
	if err != nil {
		return nil, err
	}

	c := mysql.Config{
		DBName:    os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Addr:      os.Getenv("DB_ADDR"),
		Net:       os.Getenv("DB_NET"),
		ParseTime: os.Getenv("DB_PARSE_TIME") == "true",
		Collation: os.Getenv("DB_COLLATION"),
		Loc:       jst,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
