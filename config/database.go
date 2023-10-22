package config

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var once sync.Once

func InitDB() {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Cfg.Database.Host, Cfg.Database.Port, Cfg.Database.Username, Cfg.Database.Password, Cfg.Database.Dbname)

	db, err := sql.Open("postgres", info)

	if err != nil {
		panic(err)
	}

	once.Do(func() {
		DB = db
		fmt.Println("Successfully Connect To Database!")
	})
}
