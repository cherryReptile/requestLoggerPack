package pkg

import (
	"fmt"
	"sync"
)

type GoLogger struct {
	Connect *connect
	Wg      *sync.WaitGroup
}

func NewGoLogger(url, token string, wg *sync.WaitGroup) *GoLogger {
	return &GoLogger{
		Connect: NewConnect(url, token),
		Wg:      wg,
	}
}

func (l *GoLogger) LogINFO(data any, errCh chan error) {
	l.goLog(data, errCh, "info")
}

func (l *GoLogger) logWARN(data any, errCh chan error) {
	l.goLog(data, errCh, "warn")
}

func (l *GoLogger) logERROR(data any, errCh chan error) {
	l.goLog(data, errCh, "error")
}

func (l *GoLogger) logFATAL(data any, errCh chan error) {
	l.goLog(data, errCh, "fatal")
}

func (l *GoLogger) goLog(data any, errCh chan error, level string) {
	if l.Wg != nil {
		defer l.Wg.Done()
	}

	if res, err := l.Connect.log(data, level); err != nil {
		errCh <- err
	} else if res.StatusCode != 201 {
		errCh <- fmt.Errorf("failed to create log\nresponse: %v", res)
	}
}
