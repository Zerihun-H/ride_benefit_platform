package rest

import (
	"net/http"
	partner "rideBenefit/internal/module/partner"

	"github.com/julienschmidt/httprouter"
)

// PartnerHandler contains the function of handler for domain Partner
type PartnerHandler interface {
	GetPartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddPartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdatePartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeletePartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type partnerHandler struct {
	PartnerCase partner.Usecase
}

// PartnerInit is to initialize the rest handler for domain Partner
func PartnerInit(PartnerCase partner.Usecase) PartnerHandler {
	return &partnerHandler{
		PartnerCase,
	}
}

func (dh *partnerHandler) GetPartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (dh *partnerHandler) AddPartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *partnerHandler) UpdatePartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (dh *partnerHandler) DeletePartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
