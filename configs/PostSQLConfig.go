package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn
var connErr error

var host string
var database string
var user string
var password string
var url string

func init() {
	host = os.Getenv("S4C_HOST")
	database = os.Getenv("S4C_DATABASE")
	user = os.Getenv("S4C_USERNAME")
	password = os.Getenv("S4C_PASSWORD")
	fmt.Printf("Host %s\n", host)
	fmt.Printf("Database %s\n", database)
	url = "postgresql://" + user + ":" + password + "@" + host + ":5432/" + database
	Conn, connErr = pgx.Connect(context.Background(), url)
	if connErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", connErr)
		os.Exit(1)
	}
}
