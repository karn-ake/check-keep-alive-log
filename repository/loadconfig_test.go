package repository

import (
	"fmt"
	"testing"

	"../entity"
)

func TestLoadConfig(t *testing.T) {
	var repo Repository = NewGetRepository()
	var cfg entity.FileConfig
	repo.LoadConfig(&cfg)
	fmt.Printf("%v", cfg.BlpFile)
	fmt.Printf("%v", cfg.ClvFile)
}
