package model

import "errors"

const (
	RoleAdmin      = "admin"
	RoleSuperAdmin = "superadmin"
	RolePartner    = "partner"
)

var ErrInvalidDriverID = errors.New("invalid driver ID")
var ErrInvalidRequestBody = errors.New("invalid request body")
