package model

import "errors"

const (
	RoleAdmin      = "admin"
	RoleSuperAdmin = "superadmin"
	RolePartner    = "partner"
)

var ErrInvalidRequestBody = errors.New("invalid request body")
var ErrInvalidEmployeeID = errors.New("invalid employee ID")
var ErrInvalidUserID = errors.New("invalid user ID")
var ErrInvalidPartnerID = errors.New("invalid partner ID")
var ErrInvalidRelativeID = errors.New("invalid relative ID")
var ErrInvalidServiceID = errors.New("invalid service ID")
