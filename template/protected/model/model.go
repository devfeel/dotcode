package model
import "github.com/devfeel/dotcode/template"

func init(){
	template.RegisterTemplate(new(Model))
}


type Model struct{}

func(t *Model) Path() string{
	return "protected/model"
}

func (t *Model) FileName() string{
	return "DemoInfo.go"
}

func (t *Model) String() string {
	return `package model

type DemoInfo struct {
	ID       int
	DemoID   int
	DemoName string
}`
}
