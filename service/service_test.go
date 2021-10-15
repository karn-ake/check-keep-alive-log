package service

import (
	"fmt"
	"testing"
)

func TestFlogTime(t *testing.T) {
	var time Service = NewGetService()
	fmt.Println(time.GetLocalLogTime())
	fmt.Println(time.GetTimes())
	fmt.Println(time.GetKALocalLogTime())
	fmt.Println(time.GetKATimes())
	fmt.Println(time.GetKSLocalLogTime())
	fmt.Println(time.GetKSTimes())
	fmt.Println(time.GetKTLocalLogTime())
	fmt.Println(time.GetKTTimes())
	fmt.Println(time.GetMFCLocalLogTime())
	fmt.Println(time.GetMFCTimes())
	fmt.Println(time.GetSCBLocalLogTime())
	fmt.Println(time.GetSCBTimes())
}
