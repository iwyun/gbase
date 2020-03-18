package utils

import (
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

const (
	headerXForwardedFor = "X-Forwarded-For"
	headerXRealIP       = "X-Real-IP"
)

func GetRealIP(c *gin.Context) (ip string) {
	ra := c.Request.RemoteAddr
	if ip := c.Request.Header.Get(headerXForwardedFor); ip != "" {
		ra = strings.Split(ip, ", ")[0]
	} else if ip := c.Request.Header.Get(headerXRealIP); ip != "" {
		ra = ip
	} else {
		ra, _, _ = net.SplitHostPort(ra)
	}
	return ra
}
