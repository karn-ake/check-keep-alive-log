package main

import (
	"fmt"
	"net/http"

	"./controller"
	"./router"
	"github.com/karn-ake/go_cns/controller"
)

var (
	controllers controller.GetStatusService = controller.NewGetStatusService()
	muxRouter   router.MuxRouter            = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"
	muxRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and Running")
	})
	muxRouter.GET("/api/blp", controller.CheckBlpStatus())

	muxRouter.SERV(port)
}
