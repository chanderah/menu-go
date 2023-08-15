package util

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chanderah/menu-go/model"
	"github.com/jlaffaye/ftp"
)

const HOST = "ftp.chandrasa.fun:21"
const USERNAME = "chandra5"
const PASSWORD = "dhearbiznmd"
const PATH = "/public_html/go/public"

func getFtpConnection() (*ftp.ServerConn, error) {
	conn, err := ftp.Dial(HOST, ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		return nil, err
	}
	if err = conn.Login(USERNAME, PASSWORD); err != nil {
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

	entries, err := conn.List(PATH)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		log.Println(entry.Name, entry.Type, entry.Size)
	}
	return entries, nil
}

func UploadFile(fileDetails *model.FileDetails, file *os.File) (interface{}, error) {
	conn, err := getFtpConnection()
	if err != nil {
		return nil, err
	}
	defer os.Remove(fileDetails.Dest)
	defer file.Close()
	defer conn.Quit()

	if err := conn.Stor(fileDetails.Dest, file); err != nil {
		return nil, err
	}
	log.Println("File is uploaded to: " + PATH + fileDetails.Dest)
	return PATH + fileDetails.Dest, nil
}

func UploadFile2(fileDetails *model.FileDetails) (interface{}, error) {
	conn, err := getFtpConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Quit()

	decoded, err := Decode64(strings.Split(fileDetails.File, ",")[1])
	if err != nil {
		return nil, err
	}

	dest := PATH + fileDetails.Dest
	if err := conn.Stor(dest, bytes.NewReader(decoded)); err != nil {
		return nil, err
	}
	return fmt.Sprint("File is uploaded to:", dest), nil
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
