package main

import (
	"errors"

	"github.com/qianlnk/goerr"
)

func productErr() goerr.Goerr {
	err := goerr.Err(errors.New("Error: goerr."))
	if err != nil {
		err.AddValue("age", 26)
		err.AddValue("age", 27) //a warning will occur，“key age exist!”
	}

	return err
}

func productNil() goerr.Goerr {
	err := goerr.Err(nil)
	if err != nil {
		err.AddValue("age", 26)
		err.AddValue("name", "qianlnk")
	}
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
	if err != nil {
		err.Stdout()
	}
}
