package validate

import "emoney.cn/createcode/template"

func init(){
	template.RegisterTemplate(new(Validate))
}


type Validate struct{}

func(t *Validate) Path() string{
	return "server/validate"
}

func (t *Validate) FileName() string{
	return "validate.go"
}

func (t *Validate) String() string {
	return `package validate

import (
	"{project}/server/contract"
	"errors"
)

func IsNilString(val string, errCode int, errMsg string) (*contract.ResponseInfo, error) {
	if val != "" {
		return contract.NewResonseInfo(), nil
	} else {
		return contract.CreateResponse(errCode, errMsg, nil), errors.New("val is nil")
	}
}`
}

