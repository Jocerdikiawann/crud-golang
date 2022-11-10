package rolerepositories

import (
	"belajar-golang-rest-api/models/roles"
	"context"
)

type RoleRepository interface {
	Create(ctx context.Context, req roles.Roles) (roles.Roles, error)
	GetRoles(ctx context.Context) ([]roles.Roles, error)
	GetRole(ctx context.Context, id uint) (roles.Roles, error)
	Update(ctx context.Context, req roles.Roles) (roles.Roles, error)
	Delete(ctx context.Context, id uint) error
}
