package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	"github.wdf.sap.corp/i332859/gotemplate-test/renderer"
)

func main() {
	templateData, err := ioutil.ReadFile("template.yaml")
	if err != nil {
		fmt.Printf("Failed to read template file: %v\n", err)
		return
	}
	content := string(templateData)

	valuesData, err := ioutil.ReadFile("values.yaml")
	if err != nil {
		fmt.Printf("Failed to read values file: %v\n", err)
		return
	}
	values := make(map[string]interface{})
	yaml.Unmarshal(valuesData, &values)

	renderer.RenderGoTemplate(content, values)
}
