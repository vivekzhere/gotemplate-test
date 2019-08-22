package main

import (
	"fmt"
	"io/ioutil"

	"github.wdf.sap.corp/i332859/gotemplate-test/gotemplates"
)

func main() {
	values := gotemplates.NewValues()
	values.Set(".instance.metadata.name", "hello")
	values.Set(".binding.metadata.name", "world")

	data, err := ioutil.ReadFile("template.yaml")
	if err != nil {
		fmt.Printf("Failed to read template file: %v\n", err)
		return
	}

	content := string(data)
	gotemplates.TestGoTemplate(content, values)
}
