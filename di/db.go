package di

import (
	"os"
	"test-k-style-hub/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MSYQL_PORT")
	dbname := os.Getenv("MYSQL_DB_NAME")
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWD")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return nil, err
	}

	DB.AutoMigrate(
		&entity.Members{},
		&entity.Products{},
		&entity.ReviewProducts{},
		&entity.LikeReviews{},
	)

	return DB, nil
}
