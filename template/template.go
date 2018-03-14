package template


type Template interface{
	String() string
	Path() string
	FileName() string
}

var GlobalTemplates []Template

func RegisterTemplate(t Template){
	GlobalTemplates = append(GlobalTemplates, t)
}

