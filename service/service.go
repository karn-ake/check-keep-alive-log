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
	GetLocalLogTime() string
	GetTimes() entity.AllTime
	GetKALocalLogTime() string
	GetKATimes() entity.AllTime
	GetKSLocalLogTime() string
	GetKSTimes() entity.AllTime
	GetKTLocalLogTime() string
	GetKTTimes() entity.AllTime
	GetMFCLocalLogTime() string
	GetMFCTimes() entity.AllTime
	GetSCBLocalLogTime() string
	GetSCBTimes() entity.AllTime
	GetAldnLocalLogTime() string
	GetAldnTimes() entity.AllTime
	GetInsLocalLogTime() string
	GetInsTimes() entity.AllTime
}

type getService struct{}

var (
	repo repository.Repository = repository.NewGetRepository()
)

func NewGetService() Service {
	return &getService{}
}

func (s *getService) RevFile(fName string) *[]string {
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

func (s *getService) GetLocalLogTime() string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := s.RevFile(file.BlpFile)
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
	return log
}

func (s *getService) GetTimes() (e entity.AllTime) {
	locallogtime := s.GetLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, locallogtime)
	e.LogTime = lt.Add(time.Hour * 7)
	e.SystemTime = time.Now()
	e.DiffTime = e.SystemTime.Sub(e.LogTime)
	return e
}

func (s *getService) GetKALocalLogTime() string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := s.RevFile(file.ClvFile)
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
	return log
}

func (s *getService) GetKATimes() (e entity.AllTime) {
	locallogtime := s.GetKALocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, locallogtime)
	e.LogTime = lt.Add(time.Hour * 7)
	e.SystemTime = time.Now()
	e.DiffTime = e.SystemTime.Sub(e.LogTime)
	return e
}

func (s *getService) GetKSLocalLogTime() string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := s.RevFile(file.ClvFile)
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
	return log
}

func (s *getService) GetKSTimes() (e entity.AllTime) {
	locallogtime := s.GetKSLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, locallogtime)
	e.LogTime = lt.Add(time.Hour * 7)
	e.SystemTime = time.Now()
	e.DiffTime = e.SystemTime.Sub(e.LogTime)
	return e
}

func (s *getService) GetKTLocalLogTime() string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := s.RevFile(file.ClvFile)
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
	return log
}

func (s *getService) GetKTTimes() (e entity.AllTime) {
	locallogtime := s.GetKTLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, locallogtime)
	e.LogTime = lt.Add(time.Hour * 7)
	e.SystemTime = time.Now()
	e.DiffTime = e.SystemTime.Sub(e.LogTime)
	return e
}

func (s *getService) GetMFCLocalLogTime() string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := s.RevFile(file.ClvFile)
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
	return log
}

func (s *getService) GetMFCTimes() (e entity.AllTime) {
	locallogtime := s.GetMFCLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, locallogtime)
	e.LogTime = lt.Add(time.Hour * 7)
	e.SystemTime = time.Now()
	e.DiffTime = e.SystemTime.Sub(e.LogTime)
	return e
}

func (s *getService) GetSCBLocalLogTime() string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := s.RevFile(file.ClvFile)
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
	return log
}

func (s *getService) GetSCBTimes() (e entity.AllTime) {
	locallogtime := s.GetSCBLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, locallogtime)
	e.LogTime = lt.Add(time.Hour * 7)
	e.SystemTime = time.Now()
	e.DiffTime = e.SystemTime.Sub(e.LogTime)
	return e
}

func (s *getService) GetAldnLocalLogTime() string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := s.RevFile(file.AldnFile)
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
	return log
}

func (s *getService) GetAldnTimes() (e entity.AllTime) {
	locallogtime := s.GetAldnLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, locallogtime)
	e.LogTime = lt.Add(time.Hour * 7)
	e.SystemTime = time.Now()
	e.DiffTime = e.SystemTime.Sub(e.LogTime)
	return e
}

func (s *getService) GetInsLocalLogTime() string {
	var file entity.FileConfig
	repo.LoadConfig(&file)
	rFile := s.RevFile(file.InsFile)
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
	return log
}

func (s *getService) GetInsTimes() (e entity.AllTime) {
	locallogtime := s.GetInsLocalLogTime()
	layout := "20060102-15:04:05"
	lt, _ := time.Parse(layout, locallogtime)
	e.LogTime = lt.Add(time.Hour * 7)
	e.SystemTime = time.Now()
	e.DiffTime = e.SystemTime.Sub(e.LogTime)
	return e
}
