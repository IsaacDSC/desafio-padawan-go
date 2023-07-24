package environments

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	GODEV          bool
	PORT           string
	MYSQL_USER     string
	MYSQL_PASS     string
	MYSQL_DATABASE string
	MYSQL_HOST     string
	MYSQL_PORT     string
}

var datasetEnv *Environment

func GetEnv() *Environment {
	return datasetEnv
}

func StartEnv(path string) *Environment {
	envs, err := godotenv.Read(path)
	if err != nil {
		log.Fatal(err)
	}
	var (
		port string = "3000"
	)
	if len(envs["PORT"]) > 0 {
		port = envs["PORT"]
	}
	godev, _ := strconv.ParseBool(envs["GODEV"])
	datasetEnv = &Environment{
		GODEV:          godev,
		PORT:           port,
		MYSQL_USER:     envs["MYSQL_USER"],
		MYSQL_PASS:     envs["MYSQL_PASS"],
		MYSQL_DATABASE: envs["MYSQL_DATABASE"],
		MYSQL_HOST:     envs["MYSQL_HOST"],
		MYSQL_PORT:     envs["MYSQL_PORT"],
	}
	return datasetEnv
}
