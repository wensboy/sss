package model

import "github.com/wensboy/sss/model"

type UserDao struct {
	model.SqlMeta
	Uid      int64  `db:"uid"`
	Uname    string `db:"uname"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Phone    string `db:"phone"`
}

type RoleDao struct {
	model.SqlMeta
	Rid       int64  `db:"rid"`
	Rname     string `db:"rname"`
	Creater   int64  `db:"creater"`
	InheritId int64  `db:"inherit_id"`
}

type UserRoleDao struct {
	model.SqlMeta
	UserId  int64 `db:"user_id"`
	RoleId  int64 `db:"role_id"`
	Creater int64 `db:"creater"`
}

type ActionDao struct {
	model.SqlMeta
	Aid   int64  `db:"aid"`
	Aname string `db:"aname"`
}

type ResourceDao struct {
	model.SqlMeta
	Rid   int64  `db:"rid"`
	Rname string `db:"rname"`
}

type PermissionDao struct {
	model.SqlMeta
	Pid        int64  `db:"pid"`
	ActionId   int64  `db:"action_id"`
	Action     string `db:"action"`
	ResourceId int64  `db:"resource_id"`
	Resource   string `db:"resource"`
	Creater    int64  `db:"creater"`
}

type RolePermissionDao struct {
	model.SqlMeta
	RoleId       int64 `db:"role_id"`
	PermissionId int64 `db:"permission_id"`
	Creater      int64 `db:"creater"`
}
