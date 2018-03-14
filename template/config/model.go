package config

import "emoney.cn/createcode/template"

func init(){
	template.RegisterTemplate(new(Model))
}


type Model struct{}

func(t *Model) Path() string{
	return "config"
}

func (t *Model) FileName() string{
	return "model.go"
}

func (t *Model) String() string{
	str := `package config

import (
	"encoding/xml"
)

//配置信息
type AppConfig struct {` + "\n"
	str += `    XMLName xml.Name `+"`"+`xml:"config"`+"`\n"
	str += `    AppSets []*AppSet `+"`"+`xml:"appsettings>add"`+"`\n"
	str += `    Redises []*RedisInfo `+"`"+`xml:"redises>redis"`+"`\n"
	str += `    Databases []*DataBaseInfo `+"`"+`xml:"databases>database"`+"`\n"
	str += `    AllowIps []string `+"`"+`xml:"allowips>ip"`+"`\n"
	str += `    ConfigPath string
}

//AppSetting配置
type AppSet struct {
	Key   string `+"`"+`xml:"key,attr"`+"`\n"
str += `    Value string `+"`"+`xml:"value,attr"`+"`\n"
str += `}

//Redis信息
type RedisInfo struct {
	ID string `+"`"+`xml:"id,attr"`+"`\n"
str += `    ServerUrl string `+"`"+`xml:"serverurl,attr"`+"`\n"
str += `}

//DataBase信息
type DataBaseInfo struct {
	ID string `+"`"+`xml:"id,attr"`+"`\n"
str += `    ServerUrl string `+"`"+`xml:"serverurl,attr"`+"`\n"
str += `}`

	return str
}