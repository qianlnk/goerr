# goerr

		goerr is a library which provides a track of error. 
		often we return the err in a func, and when error occur we always 
		don't know where it occured. 
		goerr will tell you the error occur at which file, line, func.

# Install
		go get github.com/qianlnk/goerr

# Interface

`AddValue(key string, value interface{})` add some k-v Info to help debug

`Message() string` get error message

`Route() []RouteNode` get where the error occur

`Value(key string) interface{}` get the k-v you have add

`Stdout()` show fmt message for debug

# Use

```golang
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
``` 

# Result

		warn: key age exist!
		Error: Error: goerr.
		Error Route:
		/Users/xxxxxx/go/src/github.com/qianlnk/goerr/demo/demo.go 10 productErr
		/Users/xxxxxx/go/src/github.com/qianlnk/goerr/demo/demo.go 29 caller1
		/Users/xxxxxx/go/src/github.com/qianlnk/goerr/demo/demo.go 33 caller2
		/Users/xxxxxx/go/src/github.com/qianlnk/goerr/demo/demo.go 37 caller3
		/Users/xxxxxx/go/src/github.com/qianlnk/goerr/demo/demo.go 41 main
		age 26

# Result explain
		from this result message, we know the err occur at file demo.go's line 10, 
		and function name is productErr, it called by caller1、 caller2、 caller3、 main