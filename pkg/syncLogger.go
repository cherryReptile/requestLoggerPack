package pkg

import (
	"net/http"
)

type Logger struct {
	Connect *connect
}

func NewLogger(uri, token string) *Logger {
	return &Logger{
		Connect: NewConnect(uri, token),
	}
}

func (l *Logger) LogINFO(data any) (*http.Response, error) {
	return l.Connect.log(data, "info")
}

func (l *Logger) LogWARN(data any) (*http.Response, error) {
	return l.Connect.log(data, "warn")
}

func (l *Logger) LogERROR(data any) (*http.Response, error) {
	return l.Connect.log(data, "error")
}

func (l *Logger) LogFATAL(data any) (*http.Response, error) {
	return l.Connect.log(data, "fatal")
}
