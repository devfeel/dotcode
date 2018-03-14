package service

import "emoney.cn/createcode/template"

func init(){
	template.RegisterTemplate(new(BaseService))
}


type BaseService struct{}

func(t *BaseService) Path() string{
	return "protected/service"
}

func (t *BaseService) FileName() string{
	return "BaseService.go"
}

func (t *BaseService) String() string {
	return `package service

import "github.com/devfeel/cache"

type BaseService struct {
	RedisCache    cache.RedisCache
}`
}