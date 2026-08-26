package rbac

import (
	"github.com/jmoiron/sqlx"
	"github.com/wensboy/sss/rbac/model"
)

type RbacRepo interface {
	InsertUser(user *model.UserDao) error
}

type rbacRepo struct {
	db *sqlx.DB
}

func NewRbacRepo() *rbacRepo {
	return &rbacRepo{}
}

func (r *rbacRepo) SetDB(db *sqlx.DB) *rbacRepo {
	r.db = db
	return r
}

func (r *rbacRepo) InsertUser(user *model.UserDao) error {
}
