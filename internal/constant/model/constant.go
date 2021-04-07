package model

import "errors"

const (
	RoleAdmin      = "admin"
	RoleSuperAdmin = "superadmin"
	RolePartner    = "partner"
)

var ErrInvalidEmployeeID = errors.New("invalid employee ID")
var ErrInvalidRequestBody = errors.New("invalid request body")
var ErrInvalidPartnerID = errors.New("invalid partner ID")
var ErrInvalidReportID = errors.New("invalid report ID")
