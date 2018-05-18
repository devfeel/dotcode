package global

import "github.com/devfeel/dotcode/template"

func init(){
	template.RegisterTemplate(new(Global))
}


type Global struct{}

func(t *Global) Path() string{
	return "global"
}

func (t *Global) FileName() string{
	return "global.go"
}

func (t *Global) String() string{
	return `package global

import (
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotlog"
	"github.com/devfeel/dottask"
	"{project}/const"
	"errors"
	"{project}/core"
)

//全局map
var GlobalItemMap *core.CMap
var DotApp *dotweb.DotWeb
var DotTask *task.TaskService
var InnerLogger dotlog.Logger

func Init(configPath string) error{
	GlobalItemMap = core.NewCMap()
	err := dotlog.StartLogService(configPath + "/dotlog.conf")
	if err != nil {
		return errors.New("log service start error => " + err.Error())
	}
	InnerLogger = dotlog.GetLogger(_const.LoggerName_Inner)
	return nil
}`
}

