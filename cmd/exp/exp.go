package main

import (
	"fmt"

	"github.com/Aoi1011/lenslocked/models"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Order struct {
	ID          int
	UserID      int
	Amount      int
	Description string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("bob1@gmail.com", "bob123")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
