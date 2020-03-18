package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/iwyun/gbase/log"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	RequestSource = "source"
	options       = "OPTIONS"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == options {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}

func LogReq() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		traceId := GetTraceIdFromGin(c)
		if traceId == "" {
			traceId = NewTraceId()
		}
		c.Request = c.Request.Clone(SetTraceIdToContext(c.Request.Context(), traceId))

		var body []byte
		switch c.ContentType() {
		case gin.MIMEJSON:
			raw, err := c.GetRawData()
			if err != nil {
				log.Errorln(traceId, err)
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(raw))
				body = raw
			}

		}
		log.Infof("[IN] %s,source: %s, ip: %s, %s, path: %s, body: %s", traceId, c.Request.Header.Get(RequestSource), c.ClientIP(), c.Request.Method, path, BytesToStr(body))
		c.Next()
		log.Infof("[OUT] %s, %v, %d", traceId, time.Now().Sub(start), c.Writer.Status())

	}
}
