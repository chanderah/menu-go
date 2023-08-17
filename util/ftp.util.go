package util

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chanderah/menu-go/model"
	"github.com/jlaffaye/ftp"
)

var FTP_PATH string

func getFtpConnection() (*ftp.ServerConn, error) {
	FTP_PATH = os.Getenv("FTP_PATH")
	FTP_HOST := os.Getenv("FTP_HOST")
	FTP_USERNAME := os.Getenv("FTP_USERNAME")
	FTP_PASSWORD := os.Getenv("FTP_PASSWORD")

	conn, err := ftp.Dial(FTP_HOST, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}
	if err = conn.Login(FTP_USERNAME, FTP_PASSWORD); err != nil {
		return nil, err
	}
	return conn, nil
}

func GetFiles() ([]*ftp.Entry, error) {
	conn, err := getFtpConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Quit()

	entries, err := conn.List(FTP_PATH)
	if err != nil {
		return nil, err
	}
	// for _, entry := range entries {
	// 	log.Println(entry.Name, entry.Type, entry.Size)
	// }
	return entries, nil
}

func UploadFile(fileDetails *model.FileDetails) (interface{}, error) {
	conn, err := getFtpConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Quit()

	decoded, err := Decode64(strings.Split(fileDetails.File, ",")[1])
	if err != nil {
		return nil, err
	}

	dest := FTP_PATH + generateFileName(fileDetails.Dest)
	if err := conn.Stor(dest, bytes.NewReader(decoded)); err != nil {
		return nil, err
	}
	return fmt.Sprint("File is uploaded to:", dest), nil
}

func generateFileName(str string) string {
	if string(str[0]) != "/" {
		str = "/" + str
	}
	indexOfLastDir:= strings.LastIndex(str, "/")
	indexOfLastDot:= strings.LastIndex(str, ".")
	if indexOfLastDot == -1 {
		return str
	}
	fileName:= str[indexOfLastDir+1:]
	fileExtension:= str[indexOfLastDot:]
	return strings.Replace(str, fileName, fmt.Sprintf("%d%s", time.Now().UnixMicro(), fileExtension), 1)
}

func RemoveFile(filePath string) error {
	conn, err := getFtpConnection()
	if err != nil {
		return err
	}
	defer conn.Quit()

	_, err = conn.Retr(filePath)
	if err != nil {
		return err
	}

	if err := conn.Delete(filePath); err != nil {
		return err
	}
	return nil
}
