package task

import "github.com/devfeel/dotcode/template"

func init(){
	template.RegisterTemplate(new(Service))
}


type Service struct{}

func(t *Service) Path() string{
	return "task"
}

func (t *Service) FileName() string{
	return "service.go"
}

func (t *Service) String() string {
	return `package task

import (
	"github.com/devfeel/dottask"
	"{project}/global"
	"{project}/task/tasks"
	"fmt"
)

func registerTask(service *task.TaskService) {
	//TODO register task to service
	service.RegisterHandler("Task_Print", tasks.Task_Print)
}

func StartTaskService(configPath string) {
	global.DotTask = task.StartNewService()

	//register all task handler
	registerTask(global.DotTask)

	//load config file
	global.DotTask.LoadConfig(configPath + "/dottask.conf")

	//start all task
	global.DotTask.StartAllTask()

	global.InnerLogger.Debug(fmt.Sprint("StartTaskService", " ", configPath, " ", global.DotTask.PrintAllCronTask()))
}

func StopTaskService() {
	if global.DotTask != nil {
		global.DotTask.StopAllTask()
	}
}`
}


