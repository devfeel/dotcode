package server

import "github.com/devfeel/dotcode/template"

func init(){
	template.RegisterTemplate(new(Router))
}


type Router struct{}

func(t *Router) Path() string{
	return "server"
}

func (t *Router) FileName() string{
	return "router.go"
}

func (t *Router) String() string {
	return `package server

import (
	"github.com/devfeel/dotweb"

	"{project}/server/handlers/test"
	"{project}/server/handlers/demo"
)

func InitRoute(server *dotweb.HttpServer) {
	g := server.Group("/test")
	g.GET("/index", test.Index)

	g = server.Group("/demo")
	g.GET("/queryinfo", demo.QueryDemoInfo)
	g.GET("/info/:demoid", demo.QueryDemoInfo)
	g.GET("/querylist", demo.QueryDemoList)
	g.GET("/add", demo.AddDemo)
}`
}


