package gotemplates

import "strings"

type ValueMap map[string]interface{}

func NewValues() ValueMap {
	return make(ValueMap)
}

func (v ValueMap) Set(key, value string) {
	path := strings.Split(key, ".")
	_set(v, value, path)
}

func _set(obj ValueMap, value string, keys []string) {
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
	var curPos ValueMap
	curObj, ok := obj[key]
	if !ok {
		curPos = NewValues()
		obj[key] = curPos
	} else {
		curPos = curObj.(ValueMap)
	}
	_set(curPos, value, keys[1:])
}
