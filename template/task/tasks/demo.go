package tasks

import "github.com/devfeel/dotcode/template"

func init(){
	template.RegisterTemplate(new(Demo))
}


type Demo struct{}

func(t *Demo) Path() string{
	return "task/tasks/demo"
}

func (t *Demo) FileName() string{
	return "demo.go"
}

func (t *Demo) String() string {
	return `package tasks

import (
	"github.com/devfeel/dottask"
	"fmt"
	"time"
)

func Task_Print(context *task.TaskContext) error {
	fmt.Println(time.Now(), "Task_Print", context.TaskID)
	return nil
}`
}


