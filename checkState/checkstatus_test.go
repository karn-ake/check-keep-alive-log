package checkState

import (
	"fmt"
	"testing"
)

func TestCheckBlpStatus(t *testing.T) {
	var status = NewGetStatusService()
	//fmt.Println(status.CheckStatus(&entity.AllTime{}))
	fmt.Println(status.CheckBlpStatus())
}

/*
	fmt.Println(status.CheckKAStatus())
	fmt.Println(status.CheckKSStatus())
	fmt.Println(status.CheckKTStatus())
	fmt.Println(status.CheckMFCStatus())
	fmt.Println(status.CheckSCBStatus())
}
*/
