package repository

import (
	"encoding/json"
	"os"

	"github.com/karn-ake/check-keep-alive-log/entity"
)

const fCfg = "D:\\Scripts\\Go\\FIX1-SET\\config\\fileconfig.json"

type Repository interface {
	LoadConfig(cfg *entity.FileConfig) error
}

type getRepository struct{}

func NewGetRepository() Repository {
	return &getRepository{}
}

func (r *getRepository) LoadConfig(cfg *entity.FileConfig) error {
	//	f, err := os.Open("F:\\Go\\FIX1-SET\\config\\fileconfig.json")
	f, err := os.Open(fCfg)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	return nil
}
