package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/f3sys/cli/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sqids/sqids-go"
	"github.com/thanhpk/randstr"
	"github.com/urfave/cli/v2"

	_ "github.com/joho/godotenv/autoload"
)

var (
	databaseURL = os.Getenv("DATABASE_URL")
)

func newDatabase(ctx context.Context) *pgx.Conn {
	db, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		slog.Default().Error("failed to create connection", "error", err)

	}

	if db.Ping(ctx) != nil {
		slog.Default().Error("failed to ping db", "error", err)

	}

	return db
}

var (
	nodeName string
	nodeType sqlc.NodeType
	// nodeIPholder string
	// nodeIP       netip.Addr
	nodeForm = huh.NewForm(
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
		// huh.NewGroup(
		// 	huh.NewInput().
		// 		Title("Node IP").
		// 		Value(&nodeIPholder).
		// 		Validate(func(s string) error {
		// 			if parsedAddr, err := netip.ParseAddr(s); err != nil {
		// 				return fmt.Errorf("invalid IP address")
		// 			} else {
		// 				nodeIP = parsedAddr
		// 				return nil
		// 			}
		// 		}),
		// ),
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

	nodes   []sqlc.Node
	nodeID  int64
	otpForm = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int64]().Value(&nodeID).OptionsFunc(func() []huh.Option[int64] {
				opts := make([]huh.Option[int64], len(nodes))
				for i, node := range nodes {
					opts[i] = huh.NewOption(node.Name, node.ID)
				}

				return opts
			}, &nodeID),
		),
	).WithTheme(huh.ThemeBase16())
)

var (
//	quantities = [MAX_GRADE][MAX_CLASS]int{
//		{40},
//		{},
//		{},
//		{},
//		{},
//		{},
//		{},
//		{34, 35, 38, 37},
//		{22, 21, 36, 35},
//		{42, 26, 27, 27},
//		{30, 28, 28, 34, 34},
//		{29, 24, 25, 34, 33},
//	}
//
// createVisitorsParams = []int32{}
// createStudentParams  = []sqlc.CreateStudentsParams{}
// visitorHeader        = []string{"id", "random", "f3sid"}
// visitorsCSV          = [][]string{}
// studentHeader        = []string{"id", "visitor_id", "grade", "class"}
// studentsCSV    = [][]string{}
// f3sidHeader = []string{"f3sid"}
// f3sidsCSV   = [][]string{}
)

const (
	MAX_GRADE = 12
	MAX_CLASS = 5
	EXTRA     = 2
)

func Sqids() (*sqids.Sqids, error) {
	sqid, err := sqids.New(sqids.Options{
		MinLength: uint8(7),
		Alphabet:  "23456789CFGHJMPQRVWX",
		Blocklist: []string{},
	})
	if err != nil {
		return nil, err
	}

	return sqid, nil
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			// {
			// 	Name:    "csv",
			// 	Aliases: []string{"c"},
			// 	Usage:   "csv operations",
			// 	Subcommands: []*cli.Command{
			// 		{
			// 			Name:    "generate",
			// 			Aliases: []string{"g"},
			// 			Usage:   "generate csv",
			// 			Action: func(cCtx *cli.Context) error {
			// 				ctx := ctx

			// 				// Initialize database connection
			// 				db := newDatabase()
			// 				defer db.Close(ctx)

			// 				// Start transaction
			// 				q, err := db.Begin(ctx)
			// 				if err != nil {
			// 					log.Fatal(err)
			// 				}
			// 				defer q.Rollback(ctx)

			// 				s := sqlc.New(q)

			// 				// Create CSV files
			// 				visitorFile, err := os.Create("visitors.csv")
			// 				if err != nil {
			// 					log.Fatal(err)
			// 				}
			// 				defer visitorFile.Close()

			// 				studentFile, err := os.Create("students.csv")
			// 				if err != nil {
			// 					log.Fatal(err)
			// 				}
			// 				defer studentFile.Close()

			// 				f3sidFile, err := os.Create("f3sids.csv")
			// 				if err != nil {
			// 					log.Fatal(err)
			// 				}
			// 				defer f3sidFile.Close()

			// 				// Create CSV writers
			// 				visitorWriter := csv.NewWriter(visitorFile)
			// 				// studentWriter := csv.NewWriter(studentFile)
			// 				f3sidWriter := csv.NewWriter(f3sidFile)

			// 				// Write headers
			// 				if err := visitorWriter.Write(visitorHeader); err != nil {
			// 					log.Fatal(err)
			// 				}
			// 				// if err := studentWriter.Write(studentHeader); err != nil {
			// 				// 	log.Fatal(err)
			// 				// }
			// 				if err := f3sidWriter.Write(f3sidHeader); err != nil {
			// 					log.Fatal(err)
			// 				}

			// 				count := 0
			// 				for grade := 0; grade < MAX_GRADE; grade++ {
			// 					for class := 0; class < MAX_CLASS; class++ {
			// 						quantity := quantities[grade][class]
			// 						if quantity > 0 {
			// 							for range quantity + 2 {
			// 								count++

			// 								random := rand.Int32N(99)

			// 								createVisitorsParams = append(createVisitorsParams, random)

			// 								createStudentParams = append(createStudentParams, sqlc.CreateStudentsParams{
			// 									VisitorID: int64(count),
			// 									Grade:     int32(grade),
			// 									Class:     int32(class),
			// 								})

			// 								sq, err := Sqids()
			// 								if err != nil {
			// 									log.Fatal(err)
			// 								}

			// 								f3sid, err := sq.Encode([]uint64{uint64(count), uint64(random)})
			// 								if err != nil {
			// 									log.Fatal(err)
			// 								}

			// 								visitorsCSV = append(visitorsCSV, []string{
			// 									strconv.Itoa(int(count)),
			// 									strconv.Itoa(int(random)),
			// 									f3sid,
			// 								})

			// 								f3sidsCSV = append(f3sidsCSV, []string{
			// 									f3sid,
			// 								})
			// 							}
			// 						}
			// 					}
			// 				}

			// 				numberOfCopy, err := s.CreateVisitors(ctx, createVisitorsParams)
			// 				if err != nil {
			// 					log.Fatal(err)
			// 				}

			// 				// numberOfCopy := int64(len(createVisitorsParams))

			// 				slog.Default().Info("created visitors", "number_of_copy", numberOfCopy)

			// 				numberOfCopy, err = s.CreateStudents(ctx, createStudentParams)
			// 				if err != nil {
			// 					log.Fatal(err)
			// 				}

			// 				// numberOfCopy = int64(len(createStudentParams))

			// 				slog.Default().Info("created students", "number_of_copy", numberOfCopy)

			// 				// Commit transaction
			// 				q.Commit(ctx)

			// 				// Write all visitors, students, and f3sids to their respective files
			// 				if err := visitorWriter.WriteAll(visitorsCSV); err != nil {
			// 					log.Fatal(err)
			// 				}
			// 				// if err := studentWriter.WriteAll(studentsCSV); err != nil {
			// 				// 	log.Fatal(err)
			// 				// }
			// 				if err := f3sidWriter.WriteAll(f3sidsCSV); err != nil {
			// 					log.Fatal(err)
			// 				}

			// 				// Flush writers before closing files
			// 				visitorWriter.Flush()
			// 				// studentWriter.Flush()
			// 				f3sidWriter.Flush()

			// 				// Check for errors in writing
			// 				if err := visitorWriter.Error(); err != nil {
			// 					log.Fatal(err)
			// 				}
			// 				// if err := studentWriter.Error(); err != nil {
			// 				// 	log.Fatal(err)
			// 				// }
			// 				if err := f3sidWriter.Error(); err != nil {
			// 					log.Fatal(err)
			// 				}

			// 				return nil
			// 			},
			// 		},
			// 		{
			// 			Name:    "format",
			// 			Aliases: []string{"f"},
			// 			Usage:   "format csv",
			// 			Action: func(cCtx *cli.Context) error {
			// 				// Read f3sids.csv
			// 				f3sidFile, err := os.Open("f3sids.csv")
			// 				if err != nil {
			// 					log.Fatal(err)
			// 				}
			// 				defer f3sidFile.Close()

			// 				f3sidReader := csv.NewReader(f3sidFile)
			// 				f3sids, err := f3sidReader.ReadAll()
			// 				if err != nil {
			// 					log.Fatal(err)
			// 				}
			// 				// Remove header
			// 				f3sids = f3sids[1:]

			// 				// Make new file
			// 				formattedFile, err := os.Create("formatted.csv")
			// 				if err != nil {
			// 					log.Fatal(err)
			// 				}
			// 				defer formattedFile.Close()

			// 				formattedWriter := csv.NewWriter(formattedFile)

			// 				formattedCSV := [][]string{}

			// 				doubledF3sids := [][]string{}
			// 				for _, sid := range f3sids {
			// 					doubledF3sids = append(doubledF3sids, sid, sid) // Append each entry twice
			// 				}

			// 				rowSize := 50
			// 				for i := range rowSize {
			// 					row := []string{}
			// 					for j := i; j < len(doubledF3sids); j += rowSize {
			// 						row = append(row, doubledF3sids[j][0])
			// 					}
			// 					formattedCSV = append(formattedCSV, row)
			// 				}

			// 				if err := formattedWriter.WriteAll(formattedCSV); err != nil {
			// 					log.Fatal(err)
			// 				}

			// 				formattedWriter.Flush()

			// 				if err := formattedWriter.Error(); err != nil {
			// 					log.Fatal(err)
			// 				}

			// 				return nil
			// 			},
			// 		},
			// 	},
			// },
			{
				Name:  "add",
				Usage: "add an item",
				Subcommands: []*cli.Command{
					{
						Name:  "node",
						Usage: "add a new node",
						Action: func(cCtx *cli.Context) error {
							ctx := context.Background()

							err := nodeForm.Run()
							if err != nil {
								log.Fatal(err)
							}

							key := randstr.String(32)

							db := newDatabase(ctx)
							defer db.Close(ctx)
							q, err := db.Begin(ctx)
							if err != nil {
								log.Fatal(err)
							}
							defer q.Rollback(ctx)
							sq := sqlc.New(q)
							node, err := sq.CreateNode(ctx, sqlc.CreateNodeParams{
								Key:  pgtype.Text{String: key, Valid: true},
								Name: nodeName,
								Type: nodeType,
							})
							if err != nil {
								log.Fatal(err)
							}
							_, err = sq.CreateBattery(ctx, node.ID)
							if err != nil {
								log.Fatal(err)
							}
							err = q.Commit(ctx)
							if err != nil {
								log.Fatal(err)
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
							ctx := context.Background()

							err := foodForm.Run()
							if err != nil {
								log.Fatal(err)
							}

							db := newDatabase(ctx)
							defer db.Close(ctx)
							q, err := db.Begin(ctx)
							if err != nil {
								log.Fatal(err)
							}
							defer q.Rollback(ctx)
							sq := sqlc.New(q)
							_, err = sq.CreateFood(ctx, sqlc.CreateFoodParams{
								Name:  foodName,
								Price: int32(foodPrice),
							})
							if err != nil {
								log.Fatal(err)
							}

							err = q.Commit(ctx)
							if err != nil {
								log.Fatal(err)
							}

							fmt.Println("Food added successfully")
							return nil
						},
					},
				},
			},
			{
				Name:  "otp",
				Usage: "generate otp",
				Action: func(cCtx *cli.Context) error {
					ctx := context.Background()

					db := newDatabase(ctx)
					defer db.Close(ctx)
					q, err := db.Begin(ctx)
					if err != nil {
						log.Fatal(err)
					}
					defer q.Rollback(ctx)
					sq := sqlc.New(q)

					nodes, err = sq.GetNodes(ctx)
					if err != nil {
						log.Fatal(err)
					}

					err = otpForm.Run()
					if err != nil {
						log.Fatal(err)
					}

					otp, err := sq.CreateOTP(ctx, sqlc.CreateOTPParams{
						Otp: pgtype.Text{
							String: randstr.String(32),
							Valid:  true,
						},
						ID: nodeID,
					})
					if err != nil {
						log.Fatal(err)
					}

					err = q.Commit(ctx)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("OTP generated successfully")
					fmt.Println("OTP:", otp.String)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
