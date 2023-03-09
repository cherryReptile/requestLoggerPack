package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"runtime"
)

type request struct {
	Line            int    `json:"calling_on_line"`
	FileWhereCalled string `json:"file_where_calling"`
	Data            any    `json:"data"`
	Type            string `json:"type"`
}

type connect struct {
	uri   string
	token string
}

func NewConnect(url, token string) *connect {
	return &connect{
		uri:   url,
		token: token,
	}
}

func (c *connect) log(data any, level string) (*http.Response, error) {
	req := new(request)

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return nil, errors.New("failed to get filename and line")
	}

	req.FileWhereCalled, req.Line = file, line
	req.Data = data
	req.Type = reflect.TypeOf(data).String()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := c.requestToLogger(body, level)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *connect) requestToLogger(body []byte, level string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", c.uri+"/"+level, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
