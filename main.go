package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/netip"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/f3sys/cli/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/thanhpk/randstr"
	"github.com/urfave/cli/v2"

	_ "github.com/joho/godotenv/autoload"
)

var (
	databaseURL = os.Getenv("DATABASE_URL")
)

func newDatabase() *pgx.Conn {
	db, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		slog.Default().Error("failed to create connection pool", "error", err)
		os.Exit(1)
	}

	if db.Ping(context.Background()) != nil {
		slog.Default().Error("failed to ping db", "error", err)
		os.Exit(1)
	}

	return db
}

var (
	nodeName     string
	nodeType     sqlc.NodeType
	nodeIPholder string
	nodeIP       netip.Addr
	nodeForm     = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Node Name").
				Value(&nodeName).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("node name cannot be empty")
					}
					return nil
				}),
		),
		huh.NewGroup(
			huh.NewSelect[sqlc.NodeType]().
				Title("Node Type").
				Options(
					huh.NewOption("Entry", sqlc.NodeTypeENTRY),
					huh.NewOption("Food Stall", sqlc.NodeTypeFOODSTALL),
					huh.NewOption("Exhibition", sqlc.NodeTypeEXHIBITION),
				).
				Value(&nodeType),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Node IP").
				Value(&nodeIPholder).
				Validate(func(s string) error {
					if parsedAddr, err := netip.ParseAddr(s); err != nil {
						return fmt.Errorf("invalid IP address")
					} else {
						nodeIP = parsedAddr
						return nil
					}
				}),
		),
	).WithTheme(huh.ThemeBase16())

	foodName        string
	foodPriceholder string
	foodPrice       int
	foodForm        = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Food Name").
				Value(&foodName).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("food name cannot be empty")
					}
					return nil
				}),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Food Price").
				Value(&foodPriceholder).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("food price cannot be empty")
					} else {
						if parsedPrice, err := strconv.Atoi(s); err != nil {
							return fmt.Errorf("invalid food price")
						} else {
							foodPrice = parsedPrice
							return nil
						}
					}
				}),
		),
	).WithTheme(huh.ThemeBase16())
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add an item",
				Subcommands: []*cli.Command{
					{
						Name:    "node",
						Aliases: []string{"n"},
						Usage:   "add a new node",
						Action: func(cCtx *cli.Context) error {
							err := nodeForm.Run()
							if err != nil {
								log.Fatal(err)
								os.Exit(1)
							}

							key := randstr.String(32)
							if err != nil {
								log.Fatal(err)
								os.Exit(1)
							}

							db := newDatabase()
							defer db.Close(context.Background())
							q, err := db.Begin(context.Background())
							if err != nil {
								log.Fatal(err)
								os.Exit(1)
							}
							defer q.Rollback(context.Background())
							sq := sqlc.New(q)
							_, err = sq.CreateNode(context.Background(), sqlc.CreateNodeParams{
								Key:  pgtype.Text{String: key, Valid: true},
								Name: nodeName,
								Ip:   &nodeIP,
								Type: nodeType,
							})
							if err != nil {
								log.Fatal(err)
								os.Exit(1)
							}
							err = q.Commit(context.Background())
							if err != nil {
								log.Fatal(err)
								os.Exit(1)
							}

							fmt.Println("Node added successfully")
							return nil
						},
					},
					{
						Name:    "food",
						Aliases: []string{"f"},
						Usage:   "add a new food",
						Action: func(cCtx *cli.Context) error {
							err := foodForm.Run()
							if err != nil {
								log.Fatal(err)
								os.Exit(1)
							}

							db := newDatabase()
							defer db.Close(context.Background())
							q, err := db.Begin(context.Background())
							if err != nil {
								log.Fatal(err)
								os.Exit(1)
							}
							defer q.Rollback(context.Background())
							sq := sqlc.New(q)
							_, err = sq.CreateFood(context.Background(), sqlc.CreateFoodParams{
								Name:  foodName,
								Price: int32(foodPrice),
							})
							if err != nil {
								log.Fatal(err)
								os.Exit(1)
							}

							err = q.Commit(context.Background())
							if err != nil {
								log.Fatal(err)
								os.Exit(1)
							}

							fmt.Println("Food added successfully")
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
