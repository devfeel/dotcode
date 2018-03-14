package repository

import "emoney.cn/createcode/template"

func init(){
	template.RegisterTemplate(new(DemoRepository))
}


type DemoRepository struct{}

func(t *DemoRepository) Path() string{
	return "protected/repository"
}

func (t *DemoRepository) FileName() string{
	return "DemoRepository.go"
}

func (t *DemoRepository) String() string {
	return `package repository

import (
	"{project}/protected"
	"{project}/protected/model"
)

type DemoRepository struct {
	BaseRepository
}

func NewDemoRepository(conf *protected.ServiceConfig) *DemoRepository {
	repository := new(DemoRepository)
	repository.Init(conf.DefaultDBConn)
	repository.InitLogger()
	return repository
}

func (repository *DemoRepository) QueryDemoInfo(demoId int) (result map[string]interface{}, err error) {
	sql := "SELECT * FROM [Demo] WITH(NOLOCK) WHERE DemoID = ? "
	return repository.FindOne(sql, demoId)
}

func (repository *DemoRepository) QueryTopDemoList(rowCount int) (result []map[string]interface{}, err error) {
	sql := "SELECT TOP 10 * FROM [Demo] WITH(NOLOCK)"
	return repository.FindList(sql)
}

func (repository *DemoRepository) InsertDemo(demo *model.DemoInfo) (n int64, err error) {
	sql := "INSERT INTO [Demo] ([DemoID], [DemoName]) VALUES(?,?)"
	return repository.Insert(sql, demo.DemoID, demo.DemoName)
}`
}