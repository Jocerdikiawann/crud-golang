package userrepositories

import (
	"belajar-golang-rest-api/models/user"
	"context"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

// Init
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Db: db,
	}
}

func (r *UserRepositoryImpl) AuthSignIn(ctx context.Context, req user.AuthSignIn) (user.User, error) {
	var user user.User
	result := r.Db.WithContext(ctx).
		Where("email = ?", req.Email).
		Preload("Roles").
		First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return user, err
	}
	
	return user, result.Error
}

func (r *UserRepositoryImpl) AuthSignUp(ctx context.Context, req user.AuthSignUp) (user.User, error) {
	createdUser := req.ToUser()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(createdUser.Password), bcrypt.DefaultCost)

	createdUser.Password = string(hashPassword[:])

	result := r.Db.
		WithContext(ctx).
		Create(&createdUser)

	for _, rolesID := range req.RolesID {
		userRole := new(user.UsersRoles)

		userRole.UserID = createdUser.ID
		userRole.RolesID = rolesID

		result = r.Db.WithContext(ctx).Create(&userRole)
	}

	return createdUser, result.Error
}

func (r *UserRepositoryImpl) GetUser(ctx context.Context, id uint) (user.User, error) {
	var user user.User
	result := r.Db.WithContext(ctx).
		Where("id = ?", id).
		Preload("Roles").
		First(&user)
	return user, result.Error
}

func (r *UserRepositoryImpl) GetUsers(ctx context.Context) ([]user.User, error) {
	var users []user.User
	result := r.Db.WithContext(ctx).
		Preload("Roles").
		Find(&users)

	return users, result.Error
}

func (r *UserRepositoryImpl) Update(ctx context.Context, id uint, req user.UserUpdate) (user.User, error) {
	userModel := req.ToUser()
	result := r.Db.Model(&userModel).
		Clauses(clause.Returning{}).
		Where("users.id = ?", req.ID)
	return userModel, result.Error
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.Db.Delete(&user.User{}, id).Error
}
