package main

import (
	"authentication-service/cmd/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")
	conn := connectToDb()
	if conn == nil {
		log.Panic("Failed to connect to Postgres, stopping service.")
	}
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", env.PORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDb() *sql.DB {
	for i := 0; i < 10; i++ {
		connection, err := openDB(env.DSN)
		if err != nil {
			fmt.Printf("Attempt %d: failed to connect to Postgres.\n", i)
			time.Sleep(time.Second * 2)
			continue
		}
		fmt.Println("Connected to Postgres.")
		return connection
	}
	return nil
}
