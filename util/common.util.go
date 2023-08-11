package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/chanderah/menu-go/model"
	"github.com/gin-gonic/gin"
)

func TypeOf(data interface{}) string {
	return reflect.TypeOf(data).String()
}

func GetPaging(paging *model.PagingInfo) {
	if IsEmpty(paging.Limit) {
		paging.Limit = 10
	}
	if IsEmpty(paging.SortField) {
		paging.SortField = "ID"
	}
	if IsEmpty(paging.SortOrder) {
		paging.SortOrder = "ASC"
	}
	if IsEmpty(paging.Field.Column) {
		paging.Field.Column = "1"
	}
	if IsEmpty(paging.Field.Value) {
		paging.Field.Value = "1"
	}
}

func StringJoin(str ...string) string {
	return strings.Join(str, " ")
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
