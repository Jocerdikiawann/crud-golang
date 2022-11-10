package test

import (
	"belajar-golang-rest-api/configs"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetLoadEnv(t *testing.T) {
	err := godotenv.Load(os.ExpandEnv("/home/cexup/Documents/Projects/RestApi/crud-golang/.env"))
	assert.Nil(t, err)
}

func TestConn(t *testing.T) {

	env := godotenv.Load(os.ExpandEnv("/home/cexup/Documents/Projects/RestApi/crud-golang/.env"))
	assert.Nil(t, env)
	//Use your auth
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	port := os.Getenv("PG_PORT")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("PG_HOST")

	_, err := configs.Connection(host, port, name, user, password)

	assert.Nil(t, err)
}
