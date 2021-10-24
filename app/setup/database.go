package setup

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"os"
)

var Db *sqlx.DB
var Err error

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	var dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",os.Getenv("DATABASE_HOST"),os.Getenv("DATABASE_PORT"),os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), "go_blog2")
	Db,Err = sqlx.Open("postgres", dsn)
}


