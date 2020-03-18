package utils

import (
	"gopkg.in/resty.v1"
	"testing"
	"time"
)

func TestGetTraceId(t *testing.T) {
	traceId := "666699999xxxxqqqq"

	_, _ = resty.SetTimeout(time.Second*2).R().SetHeader(TraceId, traceId).Get("http://")
	t.Log("end")
}

func TestNewTraceId(t *testing.T) {
	t.Log(NewTraceId())
}

func TestNewTraceIdWithExitTraceId(t *testing.T) {
	go func() {
		n := 10
		for n > 0 {
			n--
			t.Log(NewChildIdWithParentId("1234"))
		}
	}()

	go func() {
		m := 10
		for m > 0 {
			m--
			t.Log(NewChildIdWithParentId("1234"))
		}
	}()
	time.Sleep(time.Second * 2)
}

func TestNewTraceIdWithKeys(t *testing.T) {
	t.Log(NewTraceIdWithKeys("token", "roomId"))

}
