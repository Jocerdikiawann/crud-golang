package test

import (
	"belajar-golang-rest-api/configs"
	"belajar-golang-rest-api/models/roles"
	rolerepositories "belajar-golang-rest-api/repository/roleRepositories"
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func repoRole(t *testing.T) rolerepositories.RoleRepository {

	env := godotenv.Load(os.ExpandEnv("/home/cexup/Documents/Projects/RestApi/crud-golang/.env"))
	assert.Nil(t, env)
	//Use your auth
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	port := os.Getenv("PG_PORT")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("PG_HOST")

	gormDb, err := configs.Connection(host, port, name, user, password)

	assert.Nil(t, err)

	repo := rolerepositories.NewRoleRepository(gormDb)

	return repo
}

func TestCreateRole(t *testing.T) {
	repo := repoRole(t)

	actual, err := repo.Create(context.TODO(), roles.Roles{
		Name: "admin",
	})

	expect := roles.Roles{
		ID:   1,
		Name: "admin",
	}

	assert.Equal(t, expect, actual)

	assert.Nil(t, err)
}
