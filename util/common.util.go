package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func GetNewUuid() string {
	uuid := make([]byte, 4)
    if _, err := rand.Read(uuid); err != nil {
        panic(err)
    }
    return fmt.Sprintf("%X", uuid);
}

func EncryptAES(data []byte) (string, error) {
	result := make([]byte, aes.BlockSize+len(data))
	iv := result[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	block, _ := aes.NewCipher([]byte(os.Getenv("KEY")))
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(result[aes.BlockSize:], data)

	return Encode64(result), nil
}

func DecryptAES(data string) ([]byte, error) {
	encrypted, err := Decode64(data)
	if err != nil {
		return nil, err
	}
	block, _ := aes.NewCipher([]byte(os.Getenv("KEY")))
	iv := encrypted[:aes.BlockSize]
	result := encrypted[aes.BlockSize:]

	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(result, result)

	return result, nil
}

func Encode64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Decode64(encodedStr string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encodedStr)
}

func StringJoin(str ...string) string {
	return strings.Join(str, " ")
}

func WriteLog(filePath *string, data interface{}) {
	// filePath := "../log/server.log"

	// _, err := file.WriteString(fmt.Sprint(data))
	// if err != nil {

	// }
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
