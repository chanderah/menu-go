package controller

import (
	"fmt"
	"net/smtp"

	"github.com/chanderah/menu-go/response"
	"github.com/gin-gonic/gin"
)

type Email struct {
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
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

func SendMail(subject string, body string) error {
	host := "smtp.gmail.com"
	port := 587
	senderEmail := "chandrachansa@gmail.com"
	senderPassword := "ijrtzqqvhhjhyxhf"
	toEmail := []string{"annissanvll@gmail.com"}
	ccEmail := []string{"chandrachansa@gmail.com"}

	auth := smtp.PlainAuth("", senderEmail, senderPassword, host)
	smtpAddr := fmt.Sprintf("%s:%d", host, port)
	message := []byte("From: " + senderEmail + "\r\n" +
		"To: " + toEmail[0] + "\r\n" +
		"Cc: " + ccEmail[0] + "\r\n" +
		"Subject: Awesome Subject!\r\n" +
		"Pake golang nih. ihiw\r\n")

	err := smtp.SendMail(smtpAddr, auth, senderEmail, append(toEmail), []byte(message))
	if err != nil {
		return err
	}

	return nil
}

// func SendMail(subject string, message interface{}) error {
// 	const HOST = "mail.chandrasa.fun"
// 	const PORT = 587
// 	// const SENDER_NAME = "no-reply"
// 	const AUTH_EMAIL = "no-reply@chandrasa.fun"
// 	const AUTH_PASSWORD = "dhearbiznmd"

// 	headers := make(map[string]string)
// 	headers["From"] = "no-reply@chandrasa.fun"
// 	headers["To"] = to[0]
// 	headers["Subject"] = subject

// 	body := ""
// 	for k, v := range headers {
// 		body += fmt.Sprintf("%s: %s\r\n", k, v)
// 	}

// 	auth := smtp.PlainAuth("", AUTH_EMAIL, AUTH_PASSWORD, HOST)
// 	smtpAddr := fmt.Sprintf("%s:%d", HOST, PORT)

// 	err := smtp.SendMail(smtpAddr, auth, AUTH_EMAIL, append(to, cc...), []byte(body))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func SendMail(subject string, body interface{}) error {
// 	log.Println("am i called")

// 	const HOST = "mail.chandrasa.fun"
// 	const PORT = 465
// 	// const SENDER_NAME = "no-reply"
// 	const AUTH_EMAIL = "no-reply@chandrasa.fun"
// 	const AUTH_PASSWORD = "dhearbiznmd"

// 	to := []string{"chandrachansa@gmail.com"}
// 	// cc := []string{"chandra5@chandrasa.fun"}

// 	auth := smtp.PlainAuth("", AUTH_EMAIL, AUTH_PASSWORD, HOST)
// 	tlsconfig := &tls.Config{
// 		InsecureSkipVerify: false,
// 		ServerName:         HOST,
// 	}

// 	servername := fmt.Sprintf("%s:%d", HOST, PORT)
// 	conn, errs := tls.Dial("tcp", servername, tlsconfig)
// 	if errs != nil {
// 		log.Panic(errs)
// 	}

// 	c, err := smtp.NewClient(conn, HOST)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	if err = c.Auth(auth); err != nil {
// 		log.Panic(err)
// 	}
// 	if err = c.Mail(AUTH_EMAIL); err != nil {
// 		log.Panic(err)
// 	}
// 	if err = c.Rcpt("chandrachansa@gmail.com"); err != nil {
// 		log.Panic(err)
// 	}

// 	headers := make(map[string]string)
// 	headers["From"] = "no-reply@chandrasa.fun"
// 	headers["To"] = to[0]
// 	headers["Subject"] = subject

// 	message := ""
// 	for k, v := range headers {
// 		message += fmt.Sprintf("%s: %s\r\n", k, v)
// 	}
// 	message += "\r\n" + body.(string)

// 	w, err := c.Data()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	_, err = w.Write([]byte(message))
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	err = w.Close()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	c.Quit()

// 	log.Println("Mail sent!")
// 	return nil
// }
