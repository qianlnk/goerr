package goerr

import (
	"errors"
	"fmt"
	"testing"
)

func productErr() Goerr {
	return Err(errors.New("Error: goerr."), "myargs1", "myargs2")
}

func productNil() Goerr {
	err := Err(nil)
	err.AddValue("age", 26)
	err.AddValue("name", "qianlnk")
	return err
}

func caller1() Goerr {
	return productNil()
}

func caller2() Goerr {
	return caller1()
}

func caller3() Goerr {
	return caller2()
}

func TestErr(t *testing.T) {
	err := caller3()
	fmt.Println("IsErr", err.IsErr())
	fmt.Println("errmsg:", err.Message())
	for i, e := range err.Route() {
		fmt.Println(i, e.File, e.Lineno, e.FuncName)
	}

	fmt.Println("name", err.Value("name"), "age", err.Value("age"))
}
