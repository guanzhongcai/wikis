package main

import (
	"os"
	"text/template"
)

func demo_hello()  {
	stu := struct {
		Name string
		ID   int
	}{Name: "hello", ID: 11}

	// 创建模板对象, parse关联模板
	tmpl, err := template.New("test").Parse("{{.Name}} ID is {{ .ID }}")
	if err != nil {
		panic(err)
	}

	// 渲染stu为动态数据, 标准输出到终端
	err = tmpl.Execute(os.Stdout, stu)
	if err != nil {
		panic(err)
	}
}

func main() {

	demo_person()
}

// output
// hello ID is 1
