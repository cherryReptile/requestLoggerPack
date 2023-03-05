package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type Logger struct {
	uri   string
	token string
}

func NewLogger(uri, token string) *Logger {
	return &Logger{
		uri:   uri,
		token: token,
	}
}

func (l *Logger) LogINFO(data any) (*http.Response, error) {
	return l.log(data, "info")
}

func (l *Logger) LogWARN(data any) (*http.Response, error) {
	return l.log(data, "warn")
}

func (l *Logger) LogERROR(data any) (*http.Response, error) {
	return l.log(data, "error")
}

func (l *Logger) LogFATAL(data any) (*http.Response, error) {
	return l.log(data, "fatal")
}

func (l *Logger) log(data any, level string) (*http.Response, error) {
	var request struct {
		Line            int    `json:"calling_on_line"`
		FileWhereCalled string `json:"file_where_calling"`
		Data            any    `json:"data"`
		Type            string `json:"type"`
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return nil, errors.New("failed to get filename and line")
	}

	request.FileWhereCalled, request.Line = file, line
	request.Data = data
	request.Type = reflect.TypeOf(data).String()

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	res, err := l.requestToLogger(body, level)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (l *Logger) requestToLogger(body []byte, level string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", l.uri+"/"+level, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+l.token)

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
