# dotcode
code generater base on dotweb &amp; dotlog &amp; cache &amp; mapper &amp; database

# how to use
* download dotcode.exe
* run cmd like this: dotcode.exe -project={package name} -path={code path}

if run success, it will print:
```
{code path}/{package name}/main.go
{code path}/{package name}/config/config.go
{code path}/{package name}/config/model.go
{code path}/{package name}/const/const.go
{code path}/{package name}/core/cmap.go
{code path}/{package name}/core/exception/exception.go
{code path}/{package name}/document/demo.MD
{code path}/{package name}/global/global.go
{code path}/{package name}/protected/config.go
{code path}/{package name}/protected/service.go
{code path}/{package name}/protected/model/DemoInfo.go
{code path}/{package name}/protected/repository/BaseRepository.go
{code path}/{package name}/protected/repository/DemoRepository.go
{code path}/{package name}/protected/service/BaseService.go
{code path}/{package name}/protected/service/demo/DemoService.go
{code path}/{package name}/server/router.go
{code path}/{package name}/server/server.go
{code path}/{package name}/server/contract/ResponseInfo.go
{code path}/{package name}/server/handlers/test/default.go
{code path}/{package name}/server/handlers/demo/demo.go
{code path}/{package name}/server/validate/validate.go
{code path}/{package name}/util/file/file.go
{code path}/{package name}/util/json/json.go
{code path}/{package name}/util/http/http.go
{code path}/{package name}/protected/version.yaml
{code path}/{package name}/protected/config/develop/app.conf
{code path}/{package name}/protected/config/develop/dotlog.conf
{code path}/{package name}/protected/config/develop/dotweb.conf
```
