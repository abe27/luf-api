package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	// initial database
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=%s", os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPORT"), os.Getenv("SSLMODE"), os.Getenv("TZNAME"))
	if len(os.Getenv("DBPASSWORD")) > 0 {
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=%s TimeZone=%s", os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPORT"), os.Getenv("DBPASSWORD"), os.Getenv("SSLMODE"), os.Getenv("TZNAME"))
	}

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DBUSER"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))
	//   db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	// fmt.Printf("DNS: %s\n", dns)
	configs.Store, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: false,
		SkipDefaultTransaction:                   true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbt_", // table name prefix, table for `User` would be `t_users`
			SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,  // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migration DB
	configs.SetDB()
	configs.DBSeed()
}

func main() {
	// Create config variable
	config := fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Lucky Server API Service", // add custom server header
		AppName:       "API Version 1.0",
		BodyLimit:     10 * 1024 * 1024, // this is the default limit of 10MB
	}

	app := fiber.New(config)
	// app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost, http://125.25.57.91:3000, http://125.25.57.91:4040,",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Static("/", "./public")
	routes.SetUpRouter(app)
	app.Listen(fmt.Sprintf(":%s", os.Getenv("ON_PORT")))
}
