package file

import "github.com/devfeel/dotcode/template"

func init(){
	template.RegisterTemplate(new(File))
}


type File struct{}

func(t *File) Path() string{
	return "util/file"
}

func (t *File) FileName() string{
	return "file.go"
}

func (t *File) String() string {
	return `package _file

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//get filename extensions
func GetFileExt(fileName string) string {
	if fileName == "" {
		return ""
	} else {
		index := strings.LastIndex(fileName, ".")
		if index < 0 {
			return ""
		} else {
			return string(fileName[index:])
		}
	}
}

//check filename is exist
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}`
}


