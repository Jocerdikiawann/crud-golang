package test

import (
	"belajar-golang-rest-api/configs"
	"belajar-golang-rest-api/models/user"
	userrepositories "belajar-golang-rest-api/repository/userRepositories"
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func repoUser(t *testing.T) userrepositories.UserRepository {

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

	repo := userrepositories.NewUserRepository(gormDb)

	return repo
}

func TestSignUp(t *testing.T) {
	repo := repoUser(t)

	ctx := context.TODO()

	signUp, errs := repo.AuthSignUp(ctx, user.AuthSignUp{
		Email:     "putri@gmail.com",
		Password:  "passwordnya",
		FirstName: "bacot",
		LastName:  "putri",
		Address:   "bonang",
		RolesID:   []int{1, 2},
	})

	assert.Nil(t, signUp)

	assert.Nil(t, errs)
}

func TestSignIn(t *testing.T) {
	repo := repoUser(t)

	ctx := context.TODO()

	user, err := repo.AuthSignIn(ctx, user.AuthSignIn{
		Email:    "putri@gmail.com",
		Password: "passwordnya",
	})

	t.Log(user)

	assert.Nil(t, err)
}
