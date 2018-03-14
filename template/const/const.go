package _const
import "emoney.cn/createcode/template"

func init(){
	template.RegisterTemplate(new(Const))
}


type Const struct{}

func(t *Const) Path() string{
	return "const"
}

func (t *Const) FileName() string{
	return "const.go"
}

func (t *Const) String() string{
	return `package _const

const (
	Global_ProjectName = "{project}"
	Global_Version     = "V1.0 - 20171101"
)

const (
	RedisKey_ProjectPre = "{project}"
)

const (
	LoggerName_Inner   = "DefaultLogger"
	LoggerName_Mongodb = "MongodbLogger"
	LoggerName_Service = "ServiceLogger"
	LoggerName_Repository = "RepositoryLogger"
)

const (
	ApiRetCode_Ok           = 0
	ApiRetCode_Error        = -9999
	ApiRetCode_BindError    = -9100
	ApiRetCode_ServiceError = -2001
	ApiRetMsg_Ok            = ""
)`
}
