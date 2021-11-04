package controller

import "net/http"

type Controller interface {
	CheckBlp(res http.ResponseWriter, req *http.Request)
	CheckKA(res http.ResponseWriter, req *http.Request)
	CheckKS(res http.ResponseWriter, req *http.Request)
	CheckKT(res http.ResponseWriter, req *http.Request)
	CheckMFC(res http.ResponseWriter, req *http.Request)
	CheckSCB(res http.ResponseWriter, req *http.Request)
	CheckAldn(res http.ResponseWriter, req *http.Request)
	CheckIns(res http.ResponseWriter, req *http.Request)
}
