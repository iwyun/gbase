package utils

import (
	"testing"
	"time"
)

func TestFormatLocalDateDay(t *testing.T) {
	now := time.Now()
	t.Log(now.String())
	t.Log(FormatToLocalDateDay(&now))

	utc := time.Now().UTC()
	t.Log(utc.String())
	t.Log(FormatToLocalDateDay(&utc))

	n := &time.Time{}
	n = nil
	t.Log(FormatToLocalDateDay(n))

}

func TestFormatLocalDateTime(t *testing.T) {
	now := time.Now()
	t.Log(now.String())
	t.Log(FormatToLocalDateTime(&now))

	utc := time.Now().UTC()
	t.Log(utc.String())
	t.Log(FormatToLocalDateTime(&utc))

	n := &time.Time{}
	n = nil
	t.Log(FormatToLocalDateTime(n))
}

func TestFormatLocalDateDayAndTime(t *testing.T) {
	now := time.Now()
	t.Log(now.String())
	t.Log(FormatToLocalDateDayAndTime(&now))

	utc := time.Now().UTC()
	t.Log(utc.String())
	t.Log(FormatToLocalDateDayAndTime(&utc))

	n := &time.Time{}
	n = nil
	t.Log(FormatToLocalDateDayAndTime(n))

}
