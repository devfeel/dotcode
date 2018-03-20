package develop

import "github.com/devfeel/dotcode/template"

func init(){
	template.RegisterTemplate(new(DotWebConf))
}


type DotWebConf struct{}

func(t *DotWebConf) Path() string{
	return "protected/config/develop"
}

func (t *DotWebConf) FileName() string{
	return "dotweb.conf"
}

func (t *DotWebConf) String() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<config>
<app logpath="logs/" enabledlog="true" runmode="development"/>
<server isrun="true" port="8080" />
<appset>
    <set key="set1" value="1" />
</appset>
</config>　
`
}
