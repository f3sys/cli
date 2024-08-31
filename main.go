package main

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var (
	name   string
	typeOf string
)

type User struct {
	ID       string
	Password string
}

type Node struct {
	ID     big.Int
	UserId string
	Name   string
	Type   string
	Price  int
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	schema   = os.Getenv("DB_SCHEMA")
)

func generateRandomString(n int, letters string) (string, error) {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func push(name string, typeOf string) {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	key, err := generateRandomString(32, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	if err != nil {
		log.Fatal(err)
	}

	pgTx, err := conn.Begin(context.Background())
	defer pgTx.Rollback(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var id int64
	pgTx.QueryRow(context.Background(), `insert into "nodes" ("key", "name", "type") values ($1, $2, $3) returning id`, key, name, typeOf).Scan(&id)
	_, err = pgTx.Exec(context.Background(), `insert into "batteries" ("node_id") values ($1)`, id)
	if err != nil {
		log.Fatal(err)
	}

	err = pgTx.Commit(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
	fmt.Println(key)
}

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What's the name?").
				Value(&name).
				Validate(func(s string) error {
					if s != "" {
						return nil
					} else {
						return errors.New("name empty")
					}
				}),

			huh.NewSelect[string]().
				Title("Choose the type").
				Options(
					huh.NewOption("Entry", "ENTRY"),
					huh.NewOption("Exhibition", "EXHIBITION"),
					huh.NewOption("Food Stall", "FOODSTALL"),
				).
				Value(&typeOf).
				Validate(func(s string) error {
					if s != "" {
						return nil
					} else {
						return errors.New("type empty")
					}
				}),
		),
	)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	push(name, typeOf)
}
