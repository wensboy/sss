package rbac

import "github.com/jmoiron/sqlx"

type RbacRepo interface {
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
