package rest

import "github.com/wensboy/sss/rbac"

type EchoRbacHandler interface{}

type echoRbacHandler struct {
	service rbac.RbacService
}

func NewEchoRbacHandler(service rbac.RbacService) *echoRbacHandler {
	return &echoRbacHandler{
		service: service,
	}
}
