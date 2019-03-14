package http

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"urlooker/web/g"
	//"github.com/xpharos/web/http/cookie"
	"urlooker/web/http/middleware"
	"urlooker/web/http/render"
)

func Start() {
	render.Init()
	//cookie.Init()

	r := mux.NewRouter().StrictSlash(false)
	ConfigRouter(r)

	n := negroni.New()
	n.Use(middleware.NewLogger())
	n.Use(middleware.NewRecovery())
	n.UseHandler(r)
	n.Run(g.Config.Http.Listen)
}
