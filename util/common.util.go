package util

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
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

func Encode64(str string) string {
	return b64.StdEncoding.EncodeToString([]byte(str))
}

func Decode64(encodedStr string) ([]byte, error) {
	return b64.StdEncoding.DecodeString(encodedStr)
}

func StringJoin(str ...string) string {
	return strings.Join(str, " ")
}

func DeleteFile(filePath string) {
	os.Remove(filePath)
}

func WriteLog(filePath *string, data interface{}) {
	// filePath := "../log/server.log"

	// _, err := file.WriteString(fmt.Sprint(data))
	// if err != nil {

	// }
}

func ReadFile(filePath string) (string, error) {
	body, err := os.ReadFile(filePath)
	return string(body), err
	// if err != nil {
	// 	log.Println("Failed to read file.", err)
	// } else {
	// 	log.Println(string(body))
	// 	DeleteFile(filePath)
	// }
}

func CreateFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Failed to create file.", err)
		return err
	}
	fmt.Println(file)
	return nil
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
