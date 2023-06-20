package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() {
	godotenv.Load("../../.env")
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

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Create the new database
	_, err = db.Exec("CREATE DATABASE animals")
	if err != nil {
		panic("failed to create database: " + err.Error())
	}

	fmt.Println("Database created successfully!")
}
