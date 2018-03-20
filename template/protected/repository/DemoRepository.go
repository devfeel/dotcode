package repository

import "github.com/devfeel/dotcode/template"

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

func (repository *DemoRepository) QueryDemoInfo(dest interface{}, demoId int) error {
	sql := "SELECT * FROM [Demo] WITH(NOLOCK) WHERE DemoID = ? "
	return repository.FindOne(dest, sql, demoId)
}

func (repository *DemoRepository) QueryTopDemoList(dest interface{}, rowCount int) error {
	sql := "SELECT TOP 10 * FROM [Demo] WITH(NOLOCK)"
	return repository.FindList(dest, sql)
}

func (repository *DemoRepository) InsertDemo(demo *model.DemoInfo) (n int64, err error) {
	sql := "INSERT INTO [Demo] ([DemoID], [DemoName]) VALUES(?,?)"
	return repository.Insert(sql, demo.DemoID, demo.DemoName)
}

func (repository *DemoRepository) QueryByPage(dest interface{}, skip, take int)error{
	fields := "*"
	tableName := "Demo"
	where := ""
	orderBy := "ID ASC, ID DESC"
	return repository.FindListByPage(dest, tableName, fields, where, orderBy, skip, take)
}
`
}