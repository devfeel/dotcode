package develop

import "github.com/devfeel/dotcode/template"

func init(){
	template.RegisterTemplate(new(AppConf))
}


type AppConf struct{}

func(t *AppConf) Path() string{
	return "protected/config/develop"
}

func (t *AppConf) FileName() string{
	return "app.conf"
}

func (t *AppConf) String() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<config>
<appsettings>
    <add key="ConfigPath" value="XXXXXXXX" />
</appsettings>
<allowips>
    <ip></ip>
</allowips>
<redises>
    <redis id="default" serverurl="redis://127.0.0.1:6379/0"></redis>
</redises>
<databases>
    <database id="demodb" serverurl="server=127.0.0.1;port1433;database=test;user id=sa;password=123456;encrypt=disable"></database>
</databases>
</config>
`
}
