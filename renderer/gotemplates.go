package renderer

import (
	"bytes"
	"fmt"
	"text/template"
)

// RenderGoTemplate renders a gotemplate with corresponding values
// and prints the output
func RenderGoTemplate(templateString string, values map[string]interface{}) {
	engine, err := template.New("test").Funcs(getFuncMap()).Parse(templateString)
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
