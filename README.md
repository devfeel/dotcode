# dotcode
code generater base on dotweb &amp; dotlog &amp; cache &amp; mapper &amp; database

# how to use
* download dotcode.exe
* run cmd like this: dotcode.exe -project={package name} -path={code path}

if run success, it will print:
```
{code path}/main.go
{code path}/config/config.go
{code path}/config/model.go
{code path}/const/const.go
{code path}/core/cmap.go
{code path}/core/exception/exception.go
{code path}/document/demo.MD
{code path}/global/global.go
{code path}/protected/config.go
{code path}/protected/service.go
{code path}/protected/model/DemoInfo.go
{code path}/protected/repository/BaseRepository.go
{code path}/protected/repository/DemoRepository.go
{code path}/protected/service/BaseService.go
{code path}/protected/service/demo/DemoService.go
{code path}/server/router.go
{code path}/server/server.go
{code path}/server/contract/ResponseInfo.go
{code path}/server/handlers/test/default.go
{code path}/server/handlers/demo/demo.go
{code path}/server/validate/validate.go
{code path}/util/file/file.go
{code path}/util/json/json.go
{code path}/util/http/http.go
{code path}/protected/version.yaml
{code path}/protected/config/develop/app.conf
{code path}/protected/config/develop/dotlog.conf
{code path}/protected/config/develop/dotweb.conf
```
