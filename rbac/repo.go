package rbac

import (
	"github.com/jmoiron/sqlx"
	"github.com/wensboy/sss/rbac/model"
)

type RbacRepo interface {
	InsertUser(*model.UserDao) error
	UpdateUser(*model.UserDao) error
	QueryUserByUid(int64) (model.UserDao, error)
	QueryUserByEmail(string) (model.UserDao, error)
	QueryUserByPhone(string) (model.UserDao, error)
	DeleteUser(bool, int64) error
	InsertRole(*model.RoleDao) error
	UpdateRole(*model.RoleDao) error
	QueryRoleByRid(int64) (model.RoleDao, error)
	DeleteRole(bool, int64) error
	InsertAction(*model.ActionDao) error
	UpdateAction(*model.ActionDao) error
	QueryActionByAid(int64) (model.ActionDao, error)
	DeleteAction(bool, int64) error
	InsertResource(*model.ResourceDao) error
	UpdateResource(*model.ResourceDao) error
	QueryResourceByRid(int64) (model.ResourceDao, error)
	DeleteResource(bool, int64) error
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

func (r *rbacRepo) InsertUser(dao *model.UserDao) error {
	return nil
}

func (r *rbacRepo) UpdateUser(dao *model.UserDao) error {
	return nil
}

func (r *rbacRepo) QueryUserByUid(uid int64) (dao model.UserDao, err error) {
	return model.UserDao{}, nil
}

func (r *rbacRepo) QueryUserByEmail(email string) (dao model.UserDao, err error) {
	return model.UserDao{}, nil
}

func (r *rbacRepo) QueryUserByPhone(phone string) (dao model.UserDao, err error) {
	return model.UserDao{}, nil
}

func (r *rbacRepo) DeleteUser(strict bool, uid int64) error {
	return nil
}

func (r *rbacRepo) InsertRole(dao *model.RoleDao) error {
	return nil
}

func (r *rbacRepo) UpdateRole(dao *model.RoleDao) error {
	return nil
}

func (r *rbacRepo) QueryRoleByRid(rid int64) (dao model.RoleDao, err error) {
	return model.RoleDao{}, nil
}

func (r *rbacRepo) DeleteRole(strict bool, rid int64) error {
	return nil
}

func (r *rbacRepo) InsertAction(dao *model.ActionDao) error {
	return nil
}

func (r *rbacRepo) UpdateAction(dao *model.ActionDao) error {
	return nil
}

func (r *rbacRepo) QueryActionByAid(aid int64) (dao model.ActionDao, err error) {
	return model.ActionDao{}, nil
}

func (r *rbacRepo) DeleteAction(strict bool, aid int64) error {
	return nil
}

func (r *rbacRepo) InsertResource(dao *model.ResourceDao) error {
	return nil
}

func (r *rbacRepo) UpdateResource(dao *model.ResourceDao) error {
	return nil
}

func (r *rbacRepo) QueryResourceByRid(rid int64) (dao model.ResourceDao, err error) {
	return model.ResourceDao{}, nil
}

func (r *rbacRepo) DeleteResource(strict bool, rid int64) error {
	return nil
}
