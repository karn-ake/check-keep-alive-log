package main

import (
	"fmt"
	"net/http"

	ctrl "github.com/karn-ake/check-keep-alive-log/controller"
	rtr "github.com/karn-ake/check-keep-alive-log/router"
)

var (
	controllers ctrl.Controller = ctrl.NewGetController()
	muxRouter   rtr.MuxRouter   = rtr.NewMuxRouter()
)

func main() {
	const port string = ":8080"
	muxRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and Running")
	})
	muxRouter.GET("/api/aldn", controllers.CheckAldn)
	muxRouter.GET("/api/blp", controllers.CheckBlp)
	muxRouter.GET("/api/ins", controllers.CheckIns)
	muxRouter.GET("/api/ka", controllers.CheckKA)
	muxRouter.GET("/api/ks", controllers.CheckKS)
	muxRouter.GET("/api/kt", controllers.CheckKT)
	muxRouter.GET("/api/mfc", controllers.CheckMFC)
	muxRouter.GET("/api/scb", controllers.CheckSCB)

	muxRouter.SERV(port)
}
