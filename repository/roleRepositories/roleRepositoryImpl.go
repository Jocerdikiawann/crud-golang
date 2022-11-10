package rolerepositories

import (
	"belajar-golang-rest-api/models/roles"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleRepositoryImpl struct {
	Db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{
		Db: db,
	}
}

func (r *RoleRepositoryImpl) Create(ctx context.Context, req roles.Roles) (roles.Roles, error) {
	result := r.Db.WithContext(ctx).
		Create(&req)

	return req, result.Error
}

func (r *RoleRepositoryImpl) GetRoles(ctx context.Context) ([]roles.Roles, error) {
	var roles []roles.Roles
	result := r.Db.WithContext(ctx).Find(&roles)
	return roles, result.Error
}

func (r *RoleRepositoryImpl) GetRole(ctx context.Context, id uint) (roles.Roles, error) {
	var role roles.Roles
	result := r.Db.WithContext(ctx).
		Where("roles.id = ?", id).
		First(&role)
	return role, result.Error
}

func (r *RoleRepositoryImpl) Update(ctx context.Context, req roles.Roles) (roles.Roles, error) {
	result := r.Db.Model(&req).
		Clauses(clause.Returning{}).
		Where("roles.id = ?", req.ID).
		Update("name", req.Name)
	return req, result.Error
}

func (r *RoleRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.Db.Delete(&roles.Roles{}, id).Error
}
