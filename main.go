package main

import (
	"github.com/devfeel/dotcode/template"
	_ "github.com/devfeel/dotcode/template/config"
	_ "github.com/devfeel/dotcode/template/const"
	_ "github.com/devfeel/dotcode/template/core"
	_ "github.com/devfeel/dotcode/template/core/exception"
	_ "github.com/devfeel/dotcode/template/documents"
	_ "github.com/devfeel/dotcode/template/global"
	_ "github.com/devfeel/dotcode/template/protected"
	_ "github.com/devfeel/dotcode/template/protected/model"
	_ "github.com/devfeel/dotcode/template/protected/repository"
	_ "github.com/devfeel/dotcode/template/protected/service"
	_ "github.com/devfeel/dotcode/template/protected/service/demo"
	_ "github.com/devfeel/dotcode/template/server"
	_ "github.com/devfeel/dotcode/template/server/contract"
	_ "github.com/devfeel/dotcode/template/server/handlers/test"
	_ "github.com/devfeel/dotcode/template/server/handlers/demo"
	_ "github.com/devfeel/dotcode/template/server/validate"
	_ "github.com/devfeel/dotcode/template/task"
	_ "github.com/devfeel/dotcode/template/task/tasks"
	_ "github.com/devfeel/dotcode/template/util/file"
	_ "github.com/devfeel/dotcode/template/util/json"
	_ "github.com/devfeel/dotcode/template/util/http"
	_ "github.com/devfeel/dotcode/template/resources"
	_ "github.com/devfeel/dotcode/template/resources/config/develop"

	"fmt"
	"os"
	"syscall"
	"path/filepath"
	"strings"
	"log"
	"github.com/devfeel/dotcode/const"
	"flag"
)

var ProjectName string
var CodePath string

func main(){

	parseFlag()

	fmt.Println("DotCode Current Version:", _const.Version)

	for _,v:=range template.GlobalTemplates{
		path := CodePath + "/"+v.Path()
		filename := path + "/"+v.FileName()
		fmt.Println(filename)
		err := os.MkdirAll(path, 0777)
		if err != nil {
			fmt.Println(err, "create path error", path)
			return
		}
		var mode os.FileMode
		flag := syscall.O_RDWR | syscall.O_APPEND | syscall.O_CREAT
		mode = 0666
		file, err := os.OpenFile(filename, flag, mode)
		defer file.Close()
		if err != nil {
			fmt.Println(err, "OpenFile error", filename)
			return
		}
		file.WriteString(strings.Replace(v.String(),_const.ReplaceFlag_Project, ProjectName, -1))
	}
}

func parseFlag(){
	project := flag.String("project", "", "项目名称")
	codePath := flag.String("path", "", "Code目录")
	flag.Parse()

	ProjectName = *project
	if ProjectName == "" {
		fmt.Println("no set project param")
		os.Exit(-1)
	}

	CodePath = *codePath
	if CodePath == ""{
		pathDir :=GetCurrentDirectory()
		CodePath = pathDir+ "/code"
	}
}


func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	return strings.Replace(dir, "\\", "/", -1)
}


