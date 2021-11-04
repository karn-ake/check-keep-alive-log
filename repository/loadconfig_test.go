package repository

import (
	"fmt"
	"testing"

	"../entity"
)

func TestLoadConfig(t *testing.T) {
	var repo Repository = NewGetRepository()
	var cfg entity.FileConfig
	err := repo.LoadConfig(&cfg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", cfg.BlpFile)
	fmt.Printf("%v", cfg.ClvFile)
}
