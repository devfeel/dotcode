package main

import (
	"emoney.cn/createcode/template"
	_ "emoney.cn/createcode/template/config"
	_ "emoney.cn/createcode/template/const"
	_ "emoney.cn/createcode/template/core"
	_ "emoney.cn/createcode/template/core/exception"
	_ "emoney.cn/createcode/template/documents"
	_ "emoney.cn/createcode/template/global"
	_ "emoney.cn/createcode/template/protected"
	_ "emoney.cn/createcode/template/protected/model"
	_ "emoney.cn/createcode/template/protected/repository"
	_ "emoney.cn/createcode/template/protected/service"
	_ "emoney.cn/createcode/template/protected/service/demo"
	_ "emoney.cn/createcode/template/server"
	_ "emoney.cn/createcode/template/server/contract"
	_ "emoney.cn/createcode/template/server/handlers/test"
	_ "emoney.cn/createcode/template/server/handlers/demo"
	_ "emoney.cn/createcode/template/server/validate"
	_ "emoney.cn/createcode/template/util/file"
	_ "emoney.cn/createcode/template/util/json"
	_ "emoney.cn/createcode/template/util/http"
	_ "emoney.cn/createcode/template/resources"
	_ "emoney.cn/createcode/template/resources/config/develop"

	"fmt"
	"os"
	"syscall"
	"path/filepath"
	"strings"
	"log"
	"emoney.cn/createcode/const"
	"flag"
)

var ProjectName string
var CodePath string

func main(){

	parseFlag()

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


