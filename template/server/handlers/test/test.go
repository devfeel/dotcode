package demo

import "emoney.cn/createcode/template"

func init(){
	template.RegisterTemplate(new(Test))
}


type Test struct{}

func(t *Test) Path() string{
	return "server/handlers/test"
}

func (t *Test) FileName() string{
	return "default.go"
}

func (t *Test) String() string {
	return `package test

import (
	"{project}/const"
	"github.com/devfeel/dotweb"
)

func Index(ctx dotweb.Context) error {
	ctx.ViewData().Set("version", _const.Global_Version)
	return ctx.WriteString("welcome to ", _const.Global_ProjectName)
}
`
}


