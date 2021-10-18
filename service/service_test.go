package service

import (
	"fmt"
	"testing"
)

var ser Service = NewGetService()

func TestGetLocalLogTime(t *testing.T) {
	fmt.Println(*ser.GetLocalLogTime())
}


func TestGetTime(t *testing.T) {
	fmt.Println(ser.GetTimes())
}

/*
func TestGetKALocalLogTime(t *testing.T) {
	fmt.Println(ser.GetKALocalLogTime())
}
func TestGetKATime(t *testing.T) {
	fmt.Println(ser.GetKATimes())
}
func TestGetKSLocalLogTime(t *testing.T) {
	fmt.Println(ser.GetKSLocalLogTime())
}
func TestGetKSTime(t *testing.T) {
	fmt.Println(ser.GetKSTimes())
}
func TestGetKTLocalLogTime(t *testing.T) {
	fmt.Println(ser.GetKTLocalLogTime())
}
func TestGetKTTime(t *testing.T) {
	fmt.Println(ser.GetKTTimes())
}
func TestGetMFCLocalLogTime(t *testing.T) {
	fmt.Println(ser.GetMFCLocalLogTime())
}
func TestGetMFCTime(t *testing.T) {
	fmt.Println(ser.GetMFCTimes())
}
func TestGetSCBLocalLogTime(t *testing.T) {
	fmt.Println(ser.GetSCBLocalLogTime())
}
func TestGetSCBTime(t *testing.T) {
	fmt.Println(ser.GetSCBTimes())
}
*/