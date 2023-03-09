package main

import (
	"fmt"
	"log"
	"net/http"
	"reqlog/pkg"
)

type TestStruct struct {
	TestSlice []int
}

func main() {
	logger := pkg.NewLogger("http://127.0.0.1/api/log", "token")
	var res *http.Response
	var err error
	var test struct {
		TestInt    int
		TestString string
		TestStruct TestStruct
	}
	test.TestInt = 100
	test.TestString = "test_string"
	testStruct := TestStruct{TestSlice: []int{0, 144, 156}}
	test.TestStruct = testStruct
	if res, err = logger.LogINFO(test); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
