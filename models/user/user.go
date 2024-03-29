package user

import (
	"belajar-golang-rest-api/models/roles"
)

type User struct {
	ID        int           `json:"id,omitempty" gorm:"primaryKey;<-:false"`
	Email     string        `json:"email" gorm:"not null;uniqueIndex"`
	Password  string        `json:"-" gorm:"not null"`
	FirstName string        `json:"first_name" gorm:"not null"`
	LastName  string        `json:"last_name" gorm:"not null"`
	Address   string        `json:"address" gorm:"not null"`
	RolesID   []int         `json:"roles_id" gorm:"-"`
	Roles     []roles.Roles `json:"-" gorm:"many2many:users_roles"`
}

type AuthSignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthSignUp struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	RolesID   []int  `json:"roles_id"`
}

type UserUpdate struct {
	ID        int    `json:"id,omitempty" gorm:"primaryKey;<-:false"`
	Email     string `json:"email" gorm:"not null;uniqueIndex"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Address   string `json:"address" gorm:"not null"`
}

type UsersRoles struct {
	UserID  int `json:"user_id"`
	RolesID int `json:"roles_id"`
}

func (u *UserUpdate) ToUser() User {
	return User{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Address:   u.Address,
	}
}

func (auth *AuthSignUp) ToUser() User {
	return User{
		Email:     auth.Email,
		Password:  auth.Password,
		FirstName: auth.FirstName,
		LastName:  auth.LastName,
		Address:   auth.Address,
		RolesID:   auth.RolesID,
	}
}
