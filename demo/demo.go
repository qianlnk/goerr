package main

import (
	"errors"
	"fmt"

	"github.com/qianlnk/goerr"
)

func productErr() goerr.Goerr {
	err := goerr.Err(errors.New("Error: goerr."), "myargs1", "myargs2")
	err.AddValue("age", 26)
	err.AddValue("age", 27)

	return err
}

func productNil() goerr.Goerr {
	err := goerr.Err(nil)
	err.AddValue("age", 26)
	err.AddValue("name", "qianlnk")
	return err
}

func caller1() goerr.Goerr {
	return productErr()
}

func caller2() goerr.Goerr {
	return caller1()
}

func caller3() goerr.Goerr {
	return caller2()
}

func main() {
	err := caller3()
	fmt.Println("IsErr", err.IsErr())
	fmt.Println("errmsg:", err.Message())
	for i, e := range err.Route() {
		fmt.Println(i, e.File, e.Lineno, e.FuncName)
	}

	fmt.Println("name", err.Value("name"), "age", err.Value("age"))
}
