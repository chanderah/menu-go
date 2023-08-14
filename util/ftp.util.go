package util

import (
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

const HOST = "ftp.chandrasa.fun:21"
const USERNAME = "chandra5"
const PASSWORD = "dhearbiznmd"
const PATH = "/public_html/go/public"

func getFtpConnection() (*ftp.ServerConn, error) {
	conn, err := ftp.Dial(HOST, ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err = conn.Login(USERNAME, PASSWORD); err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func GetFiles() ([]*ftp.Entry, error) {
	conn, err := getFtpConnection()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer conn.Quit()

	entries, err := conn.List(PATH)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, entry := range entries {
		log.Println(entry.Name, entry.Type, entry.Size)
	}
	return entries, nil
}

func UploadFile(dest string, filePath string) (interface{}, error) {
	conn, err := getFtpConnection()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer conn.Quit()

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	file.Close()

	if err := conn.Stor(PATH+dest, file); err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("File is uploaded to: " + PATH + dest)
	return PATH + dest, nil
}

func RemoveFile(filePath string) error {
	conn, err := getFtpConnection()
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Quit()

	_, err = conn.Retr(filePath)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := conn.Delete(filePath); err != nil {
		return err
	}
	return nil
}
