package gotemplates

import (
	"bytes"
	"fmt"
	"text/template"
)

func TestGoTemplate(content string, values map[string]interface{}) {
	engine, err := template.New("test").Funcs(getFuncMap()).Parse(content)
	if err != nil {
		fmt.Printf("Failed to create renderer: %v\n", err)
		return
	}

	buf := new(bytes.Buffer)
	err = engine.Execute(buf, values)
	if err != nil {
		fmt.Printf("Failed to get render output: %v\n", err)
		return
	}
	fmt.Println("##################################################")
	fmt.Println(buf.String())
	fmt.Println("##################################################")
}
