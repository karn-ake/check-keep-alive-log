package main

import (
	"fmt"
	"net/http"

	ctrl "./controller"
	rtr "./router"
)

var (
	controllers ctrl.GetStatusService = ctrl.NewGetStatusService()
	muxRouter   rtr.MuxRouter         = rtr.NewMuxRouter()
)

func main() {
	const port string = ":8080"
	muxRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and Running")
	})
	muxRouter.GET("/api/blp", controllers.CheckBlp)

	muxRouter.SERV(port)
}
