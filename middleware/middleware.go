package middleware

import (
	"bytes"

	"github.com/chanderah/menu-go/util"
	"github.com/gin-gonic/gin"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	data, err := util.EncryptAES(b)
	if err != nil {
		data = "ERROR"
	}
	return w.ResponseWriter.Write([]byte(data))
}

func SecurityMiddleware(c *gin.Context) {
	bw := &bodyWriter{ResponseWriter: c.Writer, body: &bytes.Buffer{}}
	c.Writer = bw
	c.Next()
	// if c.Writer.Status() >= 400 {
		// log.Println("RESPONSE: " + bw.body.String())
	// }
}

func CorsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
