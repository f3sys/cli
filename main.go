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
	"github.com/charmbracelet/huh/spinner"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
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
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	username, err := generateRandomString(32)
	if err != nil {
		log.Fatal(err)
	}

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

	var users []*User
	err = pgxscan.Select(context.Background(), pgTx, &users, `insert into "User" ("id", "password") values ($1, $2) returning id, password`, username, encoded)
	if err != nil {
		log.Fatal(err)
	}

	_, err = pgTx.Exec(context.Background(), `insert into "Node" ("userId", "name", "type", "price", "updatedAt") values ($1, $2, $3, $4, now())`, username, name, typeOf, price)
	if err != nil {
		log.Fatal(err)
	}

	err = pgTx.Commit(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users[0].ID, password)
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

		action := func() {
			push(name, typeOf, i)
		}
		err = spinner.New().Title("Pushing to Database").Action(action).Run()

		if err != nil {
			log.Fatal(err)
		}
	}
}
