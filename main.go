package main

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/matthewhartstonge/argon2"
)

var (
	name   string
	typeOf string
	price  string
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

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
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

func push(name string, typeOf string, price int) {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	password, err := generateRandomString(32)
	if err != nil {
		log.Fatal(err)
	}

	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		log.Fatal(err)
	}

	pgTx, err := conn.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var id int8
	pgTx.QueryRow(context.Background(), `insert into "nodes" ("password", "name", "type", "price") values ($1, $2, $3, $4) returning id`, encoded, name, typeOf, price).Scan(&id)

	err = pgTx.Commit(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
	fmt.Println(password)
	fmt.Println(encoded)
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

			huh.NewText().
				Title("What's the price?").
				CharLimit(400).
				Value(&price).
				Validate(func(s string) error {
					if _, err := strconv.Atoi(s); err != nil {
						return err
					} else {
						return nil
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

	if i, err := strconv.Atoi(price); err != nil {
		log.Fatal(err)
	} else {
		push(name, typeOf, i)
	}
}
