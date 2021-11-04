package service

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/karn-ake/check-keep-alive-log/entity"
	"github.com/karn-ake/check-keep-alive-log/repository"
)

type Service interface {
	RevFile(fName *string) (*[]string, error)
	GetLocalLogTime() (*string, error)
	GetTimes() (*entity.AllTime, error)
	GetKALocalLogTime() (*string, error)
	GetKATimes() (*entity.AllTime, error)
	GetKSLocalLogTime() (*string, error)
	GetKSTimes() (*entity.AllTime, error)
	GetKTLocalLogTime() (*string, error)
	GetKTTimes() (*entity.AllTime, error)
	GetMFCLocalLogTime() (*string, error)
	GetMFCTimes() (*entity.AllTime, error)
	GetSCBLocalLogTime() (*string, error)
	GetSCBTimes() (*entity.AllTime, error)
	GetAldnLocalLogTime() (*string, error)
	GetAldnTimes() (*entity.AllTime, error)
	GetInsLocalLogTime() (*string, error)
	GetInsTimes() (*entity.AllTime, error)
}

type getService struct{}

var (
	repo repository.Repository = repository.NewGetRepository()
	serv Service               = NewGetService()
	a    entity.AllTime
)

const layout = "20060102-15:04:05"

func NewGetService() Service {
	return &getService{}
}

func (*getService) RevFile(fName *string) (*[]string, error) {
	file, err := os.Open(*fName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var names []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		names = append(names, s)
	}

	for i, j := 0, len(names)-1; i < j; i, j = i+1, j-1 {
		names[i], names[j] = names[j], names[i]
	}

	return &names, nil
}

func (*getService) GetLocalLogTime() (*string, error) {
	var file entity.FileConfig
	err := repo.LoadConfig(&file)
	if err != nil {
		return nil, err
	}
	rFile, err := serv.RevFile(&file.BlpFile)
	if err != nil {
		return nil, err
	}
	var logs []string
	for i, line := range *rFile {
		if i > 100 {
			break
		}
		e := `(\d{8}-\d{2}:\d{2}:\d{2})`
		r := regexp.MustCompile(e)
		log := r.FindString(line)
		logs = append(logs, log)
	}
	log := logs[0]
	return &log, nil
}

func (*getService) GetTimes() (*entity.AllTime, error) {
	locallogtime, err := serv.GetLocalLogTime()
	if err != nil {
		return nil, err
	}
	lt, err := time.Parse(layout, *locallogtime)
	if err != nil {
		return nil, err
	}
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a, nil
}

func (*getService) GetKALocalLogTime() (*string, error) {
	var file entity.FileConfig
	err := repo.LoadConfig(&file)
	if err != nil {
		return nil, err
	}
	rFile, err := serv.RevFile(&file.ClvFile)
	if err != nil {
		return nil, err
	}
	var logs []string
	for i, line := range *rFile {
		if i > 100 {
			break
		}
		if strings.Contains(line, "KASIKORN") {
			e := `(\d{8}-\d{2}:\d{2}:\d{2})`
			r := regexp.MustCompile(e)
			log := r.FindString(line)
			logs = append(logs, log)
		}
	}
	log := logs[0]
	return &log, nil
}

func (*getService) GetKATimes() (*entity.AllTime, error) {
	locallogtime, err := serv.GetKALocalLogTime()
	if err != nil {
		return nil, err
	}
	lt, err := time.Parse(layout, *locallogtime)
	if err != nil {
		return nil, err
	}
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a, nil
}

func (*getService) GetKSLocalLogTime() (*string, error) {
	var file entity.FileConfig
	err := repo.LoadConfig(&file)
	if err != nil {
		return nil, err
	}

	rFile, err := serv.RevFile(&file.ClvFile)
	if err != nil {
		return nil, err
	}

	var logs []string
	for i, line := range *rFile {
		if i > 100 {
			break
		}
		if strings.Contains(line, "KSAMCRD") {
			e := `(\d{8}-\d{2}:\d{2}:\d{2})`
			r := regexp.MustCompile(e)
			log := r.FindString(line)
			logs = append(logs, log)
		}
	}
	log := logs[0]
	return &log, nil
}

func (*getService) GetKSTimes() (*entity.AllTime, error) {
	locallogtime, err := serv.GetKSLocalLogTime()
	if err != nil {
		return nil, err
	}

	lt, err := time.Parse(layout, *locallogtime)
	if err != nil {
		return nil, err
	}

	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a, nil
}

func (*getService) GetKTLocalLogTime() (*string, error) {
	var file entity.FileConfig
	err := repo.LoadConfig(&file)
	if err != nil {
		return nil, err
	}

	rFile, err := serv.RevFile(&file.ClvFile)
	if err != nil {
		return nil, err
	}
	var logs []string
	for i, line := range *rFile {
		if i > 100 {
			break
		}
		if strings.Contains(line, "KTAMCRD") {
			e := `(\d{8}-\d{2}:\d{2}:\d{2})`
			r := regexp.MustCompile(e)
			log := r.FindString(line)
			logs = append(logs, log)
		}
	}
	log := logs[0]
	return &log, nil
}

func (*getService) GetKTTimes() (*entity.AllTime, error) {
	locallogtime, err := serv.GetKTLocalLogTime()
	if err != nil {
		return nil, err
	}

	lt, err := time.Parse(layout, *locallogtime)
	if err != nil {
		return nil, err
	}
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a, nil
}

func (*getService) GetMFCLocalLogTime() (*string, error) {
	var file entity.FileConfig
	err := repo.LoadConfig(&file)
	if err != nil {
		return nil, err
	}

	rFile, err := serv.RevFile(&file.ClvFile)
	if err != nil {
		return nil, err
	}

	var logs []string
	for i, line := range *rFile {
		if i > 100 {
			break
		}
		if strings.Contains(line, "MFCAMCRD") {
			e := `(\d{8}-\d{2}:\d{2}:\d{2})`
			r := regexp.MustCompile(e)
			log := r.FindString(line)
			logs = append(logs, log)
		}
	}
	log := logs[0]
	return &log, nil
}

func (*getService) GetMFCTimes() (*entity.AllTime, error) {
	locallogtime, err := serv.GetMFCLocalLogTime()
	if err != nil {
		return nil, err
	}

	lt, err := time.Parse(layout, *locallogtime)
	if err != nil {
		return nil, err
	}

	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a, nil
}

func (*getService) GetSCBLocalLogTime() (*string, error) {
	var file entity.FileConfig
	err := repo.LoadConfig(&file)
	if err != nil {
		return nil, err
	}
	rFile, err := serv.RevFile(&file.ClvFile)
	if err != nil {
		return nil, err
	}

	var logs []string
	for i, line := range *rFile {
		if i > 100 {
			break
		}
		if strings.Contains(line, "SCBAMCRD") {
			e := `(\d{8}-\d{2}:\d{2}:\d{2})`
			r := regexp.MustCompile(e)
			log := r.FindString(line)
			logs = append(logs, log)
		}
	}
	log := logs[0]
	return &log, nil
}

func (*getService) GetSCBTimes() (*entity.AllTime, error) {
	locallogtime, err := serv.GetSCBLocalLogTime()
	if err != nil {
		return nil, err
	}

	lt, err := time.Parse(layout, *locallogtime)
	if err != nil {
		return nil, err
	}

	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a, nil
}

func (*getService) GetAldnLocalLogTime() (*string, error) {
	var file entity.FileConfig
	err := repo.LoadConfig(&file)
	if err != nil {
		return nil, err
	}

	rFile, err := serv.RevFile(&file.AldnFile)
	if err != nil {
		return nil, err
	}

	var logs []string
	for i, line := range *rFile {
		if i > 100 {
			break
		}
		if strings.Contains(line, "NYFIX") {
			e := `(\d{8}-\d{2}:\d{2}:\d{2})`
			r := regexp.MustCompile(e)
			log := r.FindString(line)
			logs = append(logs, log)
		}
	}
	log := logs[0]
	return &log, nil
}

func (*getService) GetAldnTimes() (*entity.AllTime, error) {
	locallogtime, err := serv.GetAldnLocalLogTime()
	if err != nil {
		return nil, err
	}

	lt, err := time.Parse(layout, *locallogtime)
	if err != nil {
		return nil, err
	}

	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a, nil
}

func (*getService) GetInsLocalLogTime() (*string, error) {
	var file entity.FileConfig
	err := repo.LoadConfig(&file)
	if err != nil {
		return nil, err
	}

	rFile, err := serv.RevFile(&file.InsFile)
	if err != nil {
		return nil, err
	}

	var logs []string
	for i, line := range *rFile {
		if i > 100 {
			break
		}
		if strings.Contains(line, "INSTINET") {
			e := `(\d{8}-\d{2}:\d{2}:\d{2})`
			r := regexp.MustCompile(e)
			log := r.FindString(line)
			logs = append(logs, log)
		}
	}
	log := logs[0]
	return &log, nil
}

func (*getService) GetInsTimes() (*entity.AllTime, error) {
	locallogtime, err := serv.GetInsLocalLogTime()
	if err != nil {
		return nil, err
	}

	lt, err := time.Parse(layout, *locallogtime)
	if err != nil {
		return nil, err
	}

	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a, nil
}
