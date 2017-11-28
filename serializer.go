package serializer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type config []string

func (c config) name() string {
	if len(c) > 1 {
		return c[0]
	}
	return ""
}

func (c config) has(key string) bool {
	for _, item := range c[1:] {
		if item == key {
			return true
		}
	}
	return false
}

func Marshal(v interface{}) ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	val := reflect.ValueOf(v).Elem()
	first := true
	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("serializer")
		var c config = strings.Split(tag, ",")
		if c.has("writeOnly") {
			continue
		}
		jsonValue, err := json.Marshal(val.Field(i).Interface())
		if err != nil {
			return nil, err
		}
		if !first {
			buffer.WriteString(", ")
		} else {
			first = false
		}
		buffer.WriteString(fmt.Sprintf("\"%s\": %s", c.name(), string(jsonValue)))
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}
