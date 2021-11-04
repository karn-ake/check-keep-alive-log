package controller

import (
	"encoding/json"
	"net/http"

	"../checkState"
)

var state checkState.GetStatusService = checkState.NewGetStatusService()



type getController struct{}

func NewGetController() Controller {
	return &getController{}
}

func (*getController) CheckBlp(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post := state.CheckBlpStatus()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckKA(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post := state.CheckKAStatus()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckKS(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post := state.CheckKSStatus()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckKT(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post := state.CheckKTStatus()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckMFC(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post := state.CheckMFCStatus()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckSCB(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post := state.CheckSCBStatus()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckAldn(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post := state.CheckAldnStatus()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckIns(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post := state.CheckInsStatus()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}
