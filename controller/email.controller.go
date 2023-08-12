package controller

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"

	"github.com/chanderah/menu-go/response"
	"github.com/gin-gonic/gin"
)

type Email struct {
	Subject string      `json:"subject" binding:"required"`
	Body    interface{} `json:"body" binding:"required"`
}

func CallSendMail(c *gin.Context) {
	var req Email
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	if err := SendMail(req.Subject, req.Body); err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.OK(c, "Mail sent!")
}

func SendMail(subject string, body interface{}) error {
	log.Println("am i called")

	const HOST = "mail.chandrasa.fun"
	const PORT = 465
	// const SENDER_NAME = "no-reply"
	const AUTH_EMAIL = "no-reply@chandrasa.fun"
	const AUTH_PASSWORD = "mypassword"

	to := []string{"chandrachansa@gmail.com"}
	// cc := []string{"chandra5@chandrasa.fun"}

	auth := smtp.PlainAuth("", AUTH_EMAIL, AUTH_PASSWORD, HOST)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         HOST,
	}

	servername := fmt.Sprintf("%s:%d", HOST, PORT)
	conn, errs := tls.Dial("tcp", servername, tlsconfig)
	if errs != nil {
		log.Panic(errs)
	}

	c, err := smtp.NewClient(conn, HOST)
	if err != nil {
		log.Panic(err)
	}
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}
	if err = c.Mail(AUTH_EMAIL); err != nil {
		log.Panic(err)
	}
	if err = c.Rcpt("chandrachansa@gmail.com"); err != nil {
		log.Panic(err)
	}

	headers := make(map[string]string)
	headers["From"] = "no-reply@chandrasa.fun"
	headers["To"] = to[0]
	headers["Subject"] = subject

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body.(string)

	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}
	c.Quit()

	log.Println("Mail sent!")
	return nil
}
