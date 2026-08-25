package rbac

type RbacService interface {
}

type rbacService struct {
	repo RbacRepo
}

func NewRbacService(repo RbacRepo) *rbacService {
	return &rbacService{
		repo: repo,
	}
}
