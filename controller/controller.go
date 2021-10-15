package controller

import (
	"time"

	"../entity"
	"../service"
)

type GetStatusService interface {
	CheckStatus(t entity.AllTime) bool
	CheckBlpStatus() entity.Alert
	CheckKAStatus() entity.Alert
	CheckKSStatus() entity.Alert
	CheckKTStatus() entity.Alert
	CheckMFCStatus() entity.Alert
	CheckSCBStatus() entity.Alert
	CheckAldnStatus() entity.Alert
	CheckInsStatus() entity.Alert
}

type statusService struct{}

/*
	check health of heartbeat that not over 2 minute
*/

var (
	serv service.Service = service.NewGetService()
)

func NewGetStatusService() GetStatusService {
	return &statusService{}
}

func (s *statusService) CheckStatus(t entity.AllTime) bool {
	st := 2 * time.Minute
	var status bool
	if t.DiffTime > st {
		status = false
	} else {
		status = true
	}
	return status
}

func (s *statusService) CheckBlpStatus() (e entity.Alert) {
	t := serv.GetTimes()
	e.Client = "BLP"
	e.Status = s.CheckStatus(t)
	return e
}

func (s *statusService) CheckKAStatus() (e entity.Alert) {
	t := serv.GetKATimes()
	e.Client = "KASIKORN"
	e.Status = s.CheckStatus(t)
	return e
}

func (s *statusService) CheckKSStatus() (e entity.Alert) {
	t := serv.GetKSTimes()
	e.Client = "KSAMCRD"
	e.Status = s.CheckStatus(t)
	return e
}

func (s *statusService) CheckKTStatus() (e entity.Alert) {
	t := serv.GetKTTimes()
	e.Client = "KTAMCRD"
	e.Status = s.CheckStatus(t)
	return e
}

func (s *statusService) CheckMFCStatus() (e entity.Alert) {
	t := serv.GetMFCTimes()
	e.Client = "MFCAMCRD"
	e.Status = s.CheckStatus(t)
	return e
}

func (s *statusService) CheckSCBStatus() (e entity.Alert) {
	t := serv.GetSCBTimes()
	e.Client = "SCBAMCRD"
	e.Status = s.CheckStatus(t)
	return e
}

func (s *statusService) CheckAldnStatus() (e entity.Alert) {
	t := serv.GetSCBTimes()
	e.Client = "NYFIX"
	e.Status = s.CheckStatus(t)
	return e
}

func (s *statusService) CheckInsStatus() (e entity.Alert) {
	t := serv.GetSCBTimes()
	e.Client = "INSTINET"
	e.Status = s.CheckStatus(t)
	return e
}