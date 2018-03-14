package contract

import "emoney.cn/createcode/template"

func init(){
	template.RegisterTemplate(new(ResponseInfo))
}


type ResponseInfo struct{}

func(t *ResponseInfo) Path() string{
	return "server/contract"
}

func (t *ResponseInfo) FileName() string{
	return "ResponseInfo.go"
}

func (t *ResponseInfo) String() string {
	return `package contract

type ResponseInfo struct {
	RetCode int
	RetMsg  string
	Message interface{}
}

func NewResonseInfo() *ResponseInfo {
	return &ResponseInfo{}
}

func CreateResponse(retCode int, retMsg string, message interface{}) *ResponseInfo {
	return &ResponseInfo{
		RetCode: retCode,
		RetMsg:  retMsg,
		Message: message,
	}
}`
}
