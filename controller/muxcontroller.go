package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/karn-ake/check-keep-alive-log/checkState"
)

var state checkState.GetStatusService = checkState.NewGetStatusService()

type getController struct{}

func NewGetController() Controller {
	return &getController{}
}

func (*getController) CheckBlp(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post, err := state.CheckBlpStatus()
	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckKA(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post, err := state.CheckKAStatus()
	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckKS(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post, err := state.CheckKSStatus()
	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckKT(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post, err := state.CheckKTStatus()
	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckMFC(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post, err := state.CheckMFCStatus()
	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckSCB(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post, err := state.CheckSCBStatus()
	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckAldn(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post, err := state.CheckAldnStatus()
	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}

func (*getController) CheckIns(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	post, err := state.CheckInsStatus()
	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&post)
}
