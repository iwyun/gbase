package utils

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

const TraceId = "traceId"
const (
	traceIdTemplate        = "%s_%s"
	childIdTemplate        = "%s_%04v"
	traceIdWithKeyTemplate = "%s_%s_%s"
	randomRange            = 1000
)

const (
	TraceModeUUID   = 1
	TraceModeZipKin = 2
)

var srcTraceId = rand.NewSource(time.Now().Unix())

func NewTraceId() string {

	return fmt.Sprintf(traceIdTemplate, TraceId, strings.ToUpper(strings.Replace(uuid.NewV4().String(), "-", "", -1)))
}

func NewChildIdWithParentId(parentId string) string {
	return fmt.Sprintf(childIdTemplate, parentId, rand.New(srcTraceId).Intn(randomRange))
}

func NewTraceIdWithKeys(key ...string) string {
	return fmt.Sprintf(traceIdWithKeyTemplate, TraceId, strings.ToUpper(strings.Replace(uuid.NewV4().String(), "-", "", -1)), strings.Join(key, "_"))
}

func GetTraceId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	traceId, ok := ctx.Value(TraceId).(string)
	if !ok {
		return ""
	}

	return traceId
}

func GetTraceIdFromGin(c *gin.Context) string {
	if c == nil {
		return ""
	}
	traceId := c.Request.Header.Get(TraceId)
	if traceId != "" {
		return traceId
	}
	return GetTraceId(c.Request.Context())
}

func SetTraceIdToContext(ctx context.Context, traceId string) context.Context {

	return context.WithValue(ctx, TraceId, traceId)
}
