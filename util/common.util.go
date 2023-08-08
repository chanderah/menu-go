package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/gin-gonic/gin"
)

func TypeOf(data interface{}) string {
	return reflect.TypeOf(data).String()
}

func IsEmpty(object interface{}) bool {
	if object == nil {
		return true
	}
	if object == "" {
		return true
	}
	if object == false {
		return true
	}
	if fmt.Sprintf("%v", object) == "0" {
		return true
	}
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		}
	}
	return false
}

func ShouldBindWithoutTag(c *gin.Context, dest interface{}) error {
	if c.ContentType() != "application/json" {
		return errors.New("'BindWithoutBindingTag' only serves for application/json not for " + c.ContentType())
	}
	buf, e := io.ReadAll(c.Request.Body)
	if e != nil {
		return e
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(buf))
	return json.Unmarshal(buf, dest)
}
