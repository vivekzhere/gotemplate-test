package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func main() {
	content := `{{- $instanceName := "" }}
	{{- with .instance.metadata.name }}
	{{- $instanceName = . }}
	{{- end }}
	{{- $bindingName := "" }}
	{{- with .binding.metadata.name }}
	{{- $bindingName = . }}
	{{- end }}
	bind:
	  response: {{ (printf "'{ \"credentials\" : %s %s }'" $instanceName $bindingName ) }}
	  state: "in progress"`

	values := newValues()
	values.set(".instance.metadata.name", "hello")
	values.set(".binding.metadata.name", "world")

	testGoTemplate(content, values)
}

func testGoTemplate(content string, values map[string]interface{}) {
	engine, err := template.New("test").Parse(content)
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

type valueMap map[string]interface{}

func newValues() valueMap {
	return make(valueMap)
}

func (v valueMap) set(key, value string) {
	path := strings.Split(key, ".")
	_set(v, value, path)
}

func _set(obj valueMap, value string, keys []string) {
	if len(keys) == 0 {
		return
	}
	key := keys[0]
	if key == "" {
		_set(obj, value, keys[1:])
		return
	}
	if len(keys) == 1 {
		obj[key] = value
		return
	}
	var curPos valueMap
	curObj, ok := obj[key]
	if !ok {
		curPos = newValues()
		obj[key] = curPos
	} else {
		curPos = curObj.(valueMap)
	}
	_set(curPos, value, keys[1:])
}
