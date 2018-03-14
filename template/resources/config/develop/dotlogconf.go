package develop

import "emoney.cn/createcode/template"

func init(){
	template.RegisterTemplate(new(DotLogConf))
}


type DotLogConf struct{}

func(t *DotLogConf) Path() string{
	return "protected/config/develop"
}

func (t *DotLogConf) FileName() string{
	return "dotlog.conf"
}

func (t *DotLogConf) String() string {
	return `<?xml version="1.0" encoding="utf-8" ?>
<config>  
  <!-- 日志组件全局配置 -->
  <global islog="True" innerlogpath="logs/" innerlogencode="gb2312"/>

  <!-- 日志组件用户自定义变量 -->
  <variable>
    <var name="LogDir" value="logs/"/>
    <var name="LogDateDir" value="logs/{year}/{month}/{day}/"/>
    <var name="MailServer" value="xxxxxx"/>
    <var name="ToMail" value="xxxxxx"/>
    <var name="MailAccount" value="xxxxxx"/>
    <var name="MailPassword" value="xxxxxx"/>
    <var name="SysName" value="DemoServer"/>
  </variable>

  <!-- 日志组件日志记录媒体 -->
  <targets>
    <file>
	<target name="f1" type="File" filename="{LogDateDir}\f1.txt" encode="gb2312"/>
    </file>
    <udp>
	<target name="u1" type="Udp" remoteip="127.0.0.1:18000" encode="gb2312"/>
    </udp>
    <email>
	<target name="e1" type="EMail" mailserver="{MailServer}" mailaccount="{MailAccount}" mailpassword="{MailPassword}" tomail="{ToMail}" subject="Devfeel.DotLog log" encode="gb2312"/>
    </email>
  </targets>

  <!-- 日志对象 -->
  <loggers>
    <logger name="DefaultLogger" configmode="classics" layout="{DateTime} - {message}" />
    <logger name="ServiceLogger" configmode="classics" layout="{DateTime} - {message}" />
    <logger name="RepositoryLogger" configmode="classics" layout="{DateTime} - {message}" />
    <logger name="DemoServiceLogger" configmode="classics" layout="{DateTime} - {message}" />
    <logger name="HandlerLogger" configmode="classics" layout="{DateTime} - {message}" />
  </loggers>
  
</config>  
`
}
