package checkState

import (
	"time"

	"github.com/karn-ake/check-keep-alive-log/entity"
	"github.com/karn-ake/check-keep-alive-log/service"
)

type GetStatusService interface {
	CheckStatus(ta *entity.AllTime) bool
	CheckBlpStatus() (*entity.Customer, error)
	CheckKAStatus() (*entity.Customer, error)
	CheckKSStatus() (*entity.Customer, error)
	CheckKTStatus() (*entity.Customer, error)
	CheckMFCStatus() (*entity.Customer, error)
	CheckSCBStatus() (*entity.Customer, error)
	CheckAldnStatus() (*entity.Customer, error)
	CheckInsStatus() (*entity.Customer, error)
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

func (*statusService) CheckBlpStatus() (*entity.Customer, error) {
	t, err := serv.GetTimes()
	if err != nil {
		return nil, err
	}
	c.Client = "BLP"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c, nil
}

func (*statusService) CheckKAStatus() (*entity.Customer, error) {
	t, err := serv.GetKATimes()
	if err != nil {
		return nil, err
	}
	c.Client = "KASIKORN"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c, nil
}

func (*statusService) CheckKSStatus() (*entity.Customer, error) {
	t, err := serv.GetKSTimes()
	if err != nil {
		return nil, err
	}
	c.Client = "KSAMCRD"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c, nil
}

func (*statusService) CheckKTStatus() (*entity.Customer, error) {
	t, err := serv.GetKTTimes()
	if err != nil {
		return nil, err
	}
	c.Client = "KTAMCRD"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c, nil
}

func (*statusService) CheckMFCStatus() (*entity.Customer, error) {
	t, err := serv.GetMFCTimes()
	if err != nil {
		return nil, err
	}
	c.Client = "MFCAMCRD"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c, nil
}

func (*statusService) CheckSCBStatus() (*entity.Customer, error) {
	t, err := serv.GetSCBTimes()
	if err != nil {
		return nil, err
	}
	c.Client = "SCBAMCRD"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c, nil
}

func (*statusService) CheckAldnStatus() (*entity.Customer, error) {
	t, err := serv.GetSCBTimes()
	if err != nil {
		return nil, err
	}
	c.Client = "NYFIX"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c, nil
}

func (*statusService) CheckInsStatus() (*entity.Customer, error) {
	t, err := serv.GetSCBTimes()
	if err != nil {
		return nil, err
	}
	c.Client = "INSTINET"
	c.Status = status.CheckStatus(t)
	c.LogTime = t.LogTime
	c.SystemTime = t.SystemTime
	c.DiffTime = t.DiffTime
	return &c, nil
}
