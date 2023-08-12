// SendMailTLS not use STARTTLS common
package controller

import (
	"crypto/tls"
	"errors"
	"net"
	"net/smtp"
	"strings"
)

func SendMailTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}

	tlsconfig := &tls.Config{ServerName: host}
	if err = validateLine(from); err != nil {
		return err
	}
	for _, recp := range to {
		if err = validateLine(recp); err != nil {
			return err
		}
	}
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	defer conn.Close()
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Hello("localhost"); err != nil {
		return err
	}
	if err = c.Auth(auth); err != nil {
		return err
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

// validateLine checks to see if a line has CR or LF as per RFC 5321
func validateLine(line string) error {
	if strings.ContainsAny(line, "\n\r") {
		return errors.New("a line must not contain CR or LF")
	}
	return nil
}
