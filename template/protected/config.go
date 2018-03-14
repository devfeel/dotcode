package protected

import "emoney.cn/createcode/template"

func init(){
	template.RegisterTemplate(new(Config))
}


type Config struct{}

func(t *Config) Path() string{
	return "protected"
}

func (t *Config) FileName() string{
	return "config.go"
}

func (t *Config) String() string {
	return `package protected

type ServiceConfig struct {
	DefaultDBConn    string
	DefaultRedisConn string
}

type RedisConfig struct {
	DefaultConnString string
}`
}
