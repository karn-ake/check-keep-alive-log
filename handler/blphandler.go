package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	ctrl "../controller"
)

type Handler interface {
	BlpHandler(w http.ResponseWriter, r *http.Request)
	KAHandler(w http.ResponseWriter, r *http.Request)
	KSHandler(w http.ResponseWriter, r *http.Request)
	KTHandler(w http.ResponseWriter, r *http.Request)
	MFCHandler(w http.ResponseWriter, r *http.Request)
	SCBHandler(w http.ResponseWriter, r *http.Request)
	AldnHandler(w http.ResponseWriter, r *http.Request)
	InsHandler(w http.ResponseWriter, r *http.Request)
}

type getHandler struct{}

var router ctrl.GetStatusService = ctrl.NewGetStatusService()

func NewGetHandler() Handler {
	return &getHandler{}
}

func (*getHandler) BlpHandler(w http.ResponseWriter, r *http.Request) {
	c := router.CheckBlpStatus()
	client, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Fprintf(w, "%s", client)
}

func (*getHandler) KAHandler(w http.ResponseWriter, r *http.Request) {
	c := router.CheckKAStatus()
	client, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Fprintf(w, "%s", client)
}

func (*getHandler) KSHandler(w http.ResponseWriter, r *http.Request) {
	c := router.CheckKSStatus()
	client, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Fprintf(w, "%s", client)
}

func (*getHandler) KTHandler(w http.ResponseWriter, r *http.Request) {
	c := router.CheckKTStatus()
	client, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Fprintf(w, "%s", client)
}

func (*getHandler) MFCHandler(w http.ResponseWriter, r *http.Request) {
	c := router.CheckMFCStatus()
	client, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Fprintf(w, "%s", client)
}

func (*getHandler) SCBHandler(w http.ResponseWriter, r *http.Request) {
	c := router.CheckSCBStatus()
	client, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Fprintf(w, "%s", client)
}

func (*getHandler) AldnHandler(w http.ResponseWriter, r *http.Request) {
	c := router.CheckAldnStatus()
	client, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Fprintf(w, "%s", client)
}

func (*getHandler) InsHandler(w http.ResponseWriter, r *http.Request) {
	c := router.CheckInsStatus()
	client, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Fprintf(w, "%s", client)
}
