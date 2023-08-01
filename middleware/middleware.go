package middleware

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.buf.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinLoggingMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{ResponseWriter: c.Writer, buf: &bytes.Buffer{}}
	c.Writer = blw
	c.Next()

	if c.Writer.Status() >= 400 {
		fmt.Println("RESPONSE: " + blw.buf.String())
	}
}
