package utils

import "time"

const (
	_day        = "2006-01-02"
	_time       = "15:04:05"
	_dayAndTime = "2006-01-02 15:04:05"
)

func FormatToLocalDay(t *time.Time) (data string) {

	if t == nil {
		return
	}

	return t.Local().Format(_day)
}

func FormatToLocalTime(t *time.Time) (data string) {
	if t == nil {
		return
	}

	return t.Local().Format(_time)
}

func FormatToLocalDayAndTime(t *time.Time) (data string) {
	if t == nil {
		return
	}

	return t.Local().Format(_dayAndTime)
}

func ParseToLocalDay(t string) (data time.Time, err error) {

	data, err = time.ParseInLocation(_day, t, time.Local)
	if err != nil {
		return
	}
	return
}

func ParseToLocalTime(t string) (data time.Time, err error) {
	data, err = time.ParseInLocation(_time, t, time.Local)
	if err != nil {
		return
	}
	return

}

func ParseToLocalDayAndTime(t string) (data time.Time, err error) {
	data, err = time.ParseInLocation(_dayAndTime, t, time.Local)
	if err != nil {
		return
	}
	return

}
