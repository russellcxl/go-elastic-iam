package db

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "github.com/lib/pq"
)

var DB *gorm.DB
var requireMigration bool

func Initialise() {
	var dbName string = os.Getenv("DB_NAME")
	var dbUser string = os.Getenv("DB_USER")
	var dbHost string = os.Getenv("DB_HOST")
	var dbPort string = os.Getenv("DB_PORT")
	var region string = os.Getenv("DB_REGION")
	var profileName string = os.Getenv("AWS_PROFILE")
	var dbEndpoint string = fmt.Sprintf("%s:%s", dbHost, dbPort)

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profileName))
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	authenticationToken, err := auth.BuildAuthToken(
		context.TODO(), dbEndpoint, region, dbUser, cfg.Credentials)
	if err != nil {
		panic("failed to create authentication token: " + err.Error())
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, authenticationToken, dbName,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect to db: " + err.Error())
	}

	fmt.Println("Connected to postgres!")

	if requireMigration {
		migrate()
	}
}
