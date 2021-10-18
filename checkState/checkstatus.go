package checkState

import (
	"time"

	"../entity"
	"../service"
)

type GetStatusService interface {
	CheckStatus(ta *entity.AllTime) bool
	CheckBlpStatus() *entity.Customer
	CheckKAStatus() *entity.Customer
	CheckKSStatus() *entity.Customer
	CheckKTStatus() *entity.Customer
	CheckMFCStatus() *entity.Customer
	CheckSCBStatus() *entity.Customer
	CheckAldnStatus() *entity.Customer
	CheckInsStatus() *entity.Customer
}

type statusService struct{}

/*
	check health of heartbeat that not over 2 minute
*/

var (
	serv   service.Service = service.NewGetService()
	status                 = NewGetStatusService()
	c      entity.Customer
)

func NewGetStatusService() GetStatusService {
	return &statusService{}
}

func (*statusService) CheckStatus(ta *entity.AllTime) bool {
	const st time.Duration = 2 * time.Minute
	var s bool
	if s = true; ta.DiffTime > st {
		s = false
	}
	return s
}

func (*statusService) CheckBlpStatus() *entity.Customer {
	t := serv.GetTimes()
	c.Client = "BLP"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c
}

func (*statusService) CheckKAStatus() *entity.Customer {
	t := serv.GetKATimes()
	c.Client = "KASIKORN"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c
}

func (*statusService) CheckKSStatus() *entity.Customer {
	t := serv.GetKSTimes()
	c.Client = "KSAMCRD"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c
}

func (*statusService) CheckKTStatus() *entity.Customer {
	t := serv.GetKTTimes()
	c.Client = "KTAMCRD"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c
}

func (*statusService) CheckMFCStatus() *entity.Customer {
	t := serv.GetMFCTimes()
	c.Client = "MFCAMCRD"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c
}

func (*statusService) CheckSCBStatus() *entity.Customer {
	t := serv.GetSCBTimes()
	c.Client = "SCBAMCRD"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c
}

func (*statusService) CheckAldnStatus() *entity.Customer {
	t := serv.GetSCBTimes()
	c.Client = "NYFIX"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c
}

func (*statusService) CheckInsStatus() *entity.Customer {
	t := serv.GetSCBTimes()
	c.Client = "INSTINET"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c
}
