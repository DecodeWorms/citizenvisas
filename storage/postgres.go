package storage

import (
	"context"
	"fmt"
	"github.com/decodeworms/citizenvisas/config"
	"github.com/jackc/pgx/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Client struct {
	conn *pgx.Conn
	db   *gorm.DB
}

func NewConnection(ctx context.Context, config config.Config) *Client {
	log.Println("connecting to DB ...")

	uri := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DatabaseHost, config.DatabasePort, config.DatabaseUsername, config.DatabasePassword, config.DatabaseName)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Panic(fmt.Sprintf("error creating postgres connection %v", err))
	}

	client := &Client{db: db}
	log.Println("connected to DB")

	return client
}
