package controller

import (
	"fmt"
	"net/smtp"
	"os"

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
	EMAIL_HOST := os.Getenv("EMAIL_HOST")
	EMAIL_PORT := os.Getenv("EMAIL_PORT")
	EMAIL_SENDER := os.Getenv("EMAIL_SENDER")
	EMAIL_PASSWORD := os.Getenv("EMAIL_PASSWORD")
	TO_EMAIL := []string{"chandrachansa@gmail.com", "chandrasarifin@gmail.com"}
	// CC_EMAIL := []string{"chandrasarifin@gmail.com"}

	auth := smtp.PlainAuth("", EMAIL_SENDER, EMAIL_PASSWORD, EMAIL_HOST)
	smtpAddr := fmt.Sprintf("%s:%s", EMAIL_HOST, EMAIL_PORT)
	message := []byte("From: " + EMAIL_SENDER + "\r\n" +
		"To: " + TO_EMAIL[0] + "\r\n" +
		"Cc: " + TO_EMAIL[1] + "\r\n" +
		"Subject: Awesome Subject!\r\n" +
		"Pake golang nih. ihiw\r\n")

	err := smtp.SendMail(smtpAddr, auth, EMAIL_SENDER, TO_EMAIL, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

// func SendMail(subject string, message interface{}) error {
// 	const EMAIL_HOST = "mail.chandrasa.fun"
// 	const EMAIL_PORT = 587
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

// 	auth := smtp.PlainAuth("", AUTH_EMAIL, AUTH_PASSWORD, EMAIL_HOST)
// 	smtpAddr := fmt.Sprintf("%s:%d", EMAIL_HOST, EMAIL_PORT)

// 	err := smtp.SendMail(smtpAddr, auth, AUTH_EMAIL, append(to, cc...), []byte(body))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func SendMail(subject string, body interface{}) error {
// 	log.Println("am i called")

// 	const EMAIL_HOST = "mail.chandrasa.fun"
// 	const EMAIL_PORT = 465
// 	// const SENDER_NAME = "no-reply"
// 	const AUTH_EMAIL = "no-reply@chandrasa.fun"
// 	const AUTH_PASSWORD = "dhearbiznmd"

// 	to := []string{"chandrachansa@gmail.com"}
// 	// cc := []string{"chandra5@chandrasa.fun"}

// 	auth := smtp.PlainAuth("", AUTH_EMAIL, AUTH_PASSWORD, EMAIL_HOST)
// 	tlsconfig := &tls.Config{
// 		InsecureSkipVerify: false,
// 		ServerName:         EMAIL_HOST,
// 	}

// 	servername := fmt.Sprintf("%s:%d", EMAIL_HOST, EMAIL_PORT)
// 	conn, errs := tls.Dial("tcp", servername, tlsconfig)
// 	if errs != nil {
// 		log.Panic(errs)
// 	}

// 	c, err := smtp.NewClient(conn, EMAIL_HOST)
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
