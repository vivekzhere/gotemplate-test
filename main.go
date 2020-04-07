package main

import (
	"fmt"
	"io/ioutil"

	"sigs.k8s.io/yaml"

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
	err = yaml.Unmarshal(valuesData, &values)
	if err != nil {
		fmt.Printf("Failed to Unmarshal values file: %v\n", err)
		return
	}

	renderer.RenderGoTemplate(content, values)
}
