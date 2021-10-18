package service

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"../entity"
	"../repository"
)

type Service interface {
	RevFile(fName string) *[]string
	GetLocalLogTime() *string
	GetTimes() *entity.AllTime
	GetKALocalLogTime() *string
	GetKATimes() *entity.AllTime
	GetKSLocalLogTime() *string
	GetKSTimes() *entity.AllTime
	GetKTLocalLogTime() *string
	GetKTTimes() *entity.AllTime
	GetMFCLocalLogTime() *string
	GetMFCTimes() *entity.AllTime
	GetSCBLocalLogTime() *string
	GetSCBTimes() *entity.AllTime
	GetAldnLocalLogTime() *string
	GetAldnTimes() *entity.AllTime
	GetInsLocalLogTime() *string
	GetInsTimes() *entity.AllTime
}

type getService struct{}

var (
	repo repository.Repository = repository.NewGetRepository()
	serv Service               = NewGetService()
	a    entity.AllTime
)

func NewGetService() Service {
	return &getService{}
}

func (*getService) RevFile(fName string) *[]string {
	file, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
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

	return &names
}

func (*getService) GetLocalLogTime() *string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := serv.RevFile(file.BlpFile)
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
	return &log
}

func (*getService) GetTimes() *entity.AllTime {
	locallogtime := serv.GetLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, *locallogtime)
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a
}

func (*getService) GetKALocalLogTime() *string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := serv.RevFile(file.ClvFile)
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
	return &log
}

func (*getService) GetKATimes() *entity.AllTime {
	locallogtime := serv.GetKALocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, *locallogtime)
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a
}

func (*getService) GetKSLocalLogTime() *string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := serv.RevFile(file.ClvFile)
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
	return &log
}

func (*getService) GetKSTimes() *entity.AllTime {
	locallogtime := serv.GetKSLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, *locallogtime)
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a
}

func (*getService) GetKTLocalLogTime() *string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := serv.RevFile(file.ClvFile)
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
	return &log
}

func (*getService) GetKTTimes() *entity.AllTime {
	locallogtime := serv.GetKTLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, *locallogtime)
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a
}

func (*getService) GetMFCLocalLogTime() *string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := serv.RevFile(file.ClvFile)
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
	return &log
}

func (*getService) GetMFCTimes() *entity.AllTime {
	locallogtime := serv.GetMFCLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, *locallogtime)
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a
}

func (*getService) GetSCBLocalLogTime() *string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := serv.RevFile(file.ClvFile)
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
	return &log
}

func (*getService) GetSCBTimes() *entity.AllTime {
	locallogtime := serv.GetSCBLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, *locallogtime)
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a
}

func (*getService) GetAldnLocalLogTime() *string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := serv.RevFile(file.AldnFile)
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
	return &log
}

func (*getService) GetAldnTimes() *entity.AllTime {
	locallogtime := serv.GetAldnLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, *locallogtime)
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a
}

func (*getService) GetInsLocalLogTime() *string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := serv.RevFile(file.InsFile)
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
	return &log
}

func (*getService) GetInsTimes() *entity.AllTime {
	locallogtime := serv.GetInsLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, *locallogtime)
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a
}
