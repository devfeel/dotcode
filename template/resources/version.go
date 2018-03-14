package resources

import "emoney.cn/createcode/template"

func init(){
	template.RegisterTemplate(new(Version))
}


type Version struct{}

func(t *Version) Path() string{
	return "protected"
}

func (t *Version) FileName() string{
	return "version.yaml"
}

func (t *Version) String() string {
	return `#version tag
version:1.0
`
}
