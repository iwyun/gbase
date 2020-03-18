package utils

import "time"

const (
	DateDay        = "2006-01-02"
	DateTime       = "15:04:05"
	DateDayAndTime = "2006-01-02 15:04:05"
)

func FormatToLocalDateDay(t *time.Time) (data string) {

	if t == nil {
		return
	}

	return t.Local().Format(DateDay)
}

func FormatToLocalDateTime(t *time.Time) (data string) {
	if t == nil {
		return
	}

	return t.Local().Format(DateTime)
}

func FormatToLocalDateDayAndTime(t *time.Time) (data string) {
	if t == nil {
		return
	}

	return t.Local().Format(DateDayAndTime)
}

func ParseToLocalDateDay(t string) (data time.Time, err error) {

	data, err = time.ParseInLocation(DateDay, t, time.Local)
	if err != nil {
		return
	}
	return
}

func ParseToLocalDateTime(t string) (data time.Time, err error) {
	data, err = time.ParseInLocation(DateTime, t, time.Local)
	if err != nil {
		return
	}
	return

}

func ParseToLocalDateDayAndTime(t string) (data time.Time, err error) {
	data, err = time.ParseInLocation(DateDayAndTime, t, time.Local)
	if err != nil {
		return
	}
	return

}
