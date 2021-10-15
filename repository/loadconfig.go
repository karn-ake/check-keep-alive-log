package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"../entity"
)

type Repository interface {
	processError(err error)
	LoadConfig(cfg *entity.FileConfig)
}

type getRepository struct{}

func NewGetRepository() Repository {
	return &getRepository{}
}

func (*getRepository) processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
func (r *getRepository) LoadConfig(cfg *entity.FileConfig) {
	f, err := os.Open("F:\\Go\\FIX1-SET\\config\\fileconfig.json")
	if err != nil {
		r.processError(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		r.processError(err)
	}
}
