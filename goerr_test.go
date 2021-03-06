package goerr

import (
	"errors"
	"fmt"
	"testing"
)

func productErr() Goerr {
	return Err(errors.New("Error: goerr."))
}

func productNil() Goerr {
	err := Err(errors.New("this a test error message."))
	if err != nil {
		err.AddValue("age", 26)
		err.AddValue("name", "qianlnk")
	}
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

	fmt.Println("errmsg:", err.Message())
	for i, e := range err.Route() {
		fmt.Println(i, e.File, e.Lineno, e.FuncName)
	}

	fmt.Println("name", err.Value("name"), "age", err.Value("age"))
}
