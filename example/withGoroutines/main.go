package main

import (
	"fmt"
	"reqlog/pkg"
	"time"
)

func main() {
	logger := pkg.NewGoLogger("http://127.0.0.1/api/log", "token", nil)
	errCh := make(chan error, 1)

	for i := 0; i <= 60; i++ {
		fmt.Println(i)
		go logger.LogINFO(i, errCh)
	}

	for {
		select {
		case err := <-errCh:
			fmt.Println(err)
		case <-time.After(time.Second * 2):
			fmt.Println("accepted signal, stopping...")
			return
		}
	}
}
