package server

import (
	"flag"
	"fmt"
	"log"
	"os"
	api "pfe/api"
	"pfe/database"
	"time"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("[WARNING]", err)
	}
}

// database connection
func DBConnection() (*gorm.DB, error) {
	// create database logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},
	)

	// database url
	url := fmt.Sprintf("host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable")

	return gorm.Open(postgres.Open(url), &gorm.Config{Logger: newLogger})
}

func enableCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// run server
func RunServer() {
	// database connection
	db, err := DBConnection()
	if err != nil {
		panic(fmt.Sprintf("[WARNING] database connection: %v", err))
	}

	// initialize casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("[WARNING] failed to initialize casbin adapter: %v", err))
	}

	// load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("[WARNING] failed to create casbin enforcer: %v", err))
	}

	// check database migration
	databaseFlag := flag.Bool("database", false, "Bool variable to create database")
	flag.Parse()

	// just create database and quit
	if *databaseFlag {
		// auto migrate tables & create root user
		database.AutoMigrateDatabase(db, enforcer)
		return
	}

	// create Gin engine
	router := gin.Default()

	// Enable CORS
	router.Use(enableCors())

	// declare API routes
	api.RoutesApis(router.Group("/api"), db, enforcer)

	// run the server
	router.Run(os.Getenv("APP_PORT"))
}
