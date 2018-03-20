package documents

import "github.com/devfeel/dotcode/template"

func init(){
	template.RegisterTemplate(new(Demo))
}


type Demo struct{}

func(t *Demo) Path() string{
	return "document"
}

func (t *Demo) FileName() string{
	return "demo.MD"
}

func (t *Demo) String() string{
	return `# demo.MD

#### 依赖项
* github.com/devfeel/dotweb
* github.com/devfeel/dotlog
* github.com/devfeel/mapper
* github.com/devfeel/middleware`
	}
