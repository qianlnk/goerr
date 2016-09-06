package goerr

import (
	"fmt"
	"runtime"
	"strings"
)

type errTracking struct {
	isErr   bool
	route   []RouteNode
	message string
	args    []interface{}
	values  map[string]interface{}
}

type RouteNode struct {
	File     string
	FuncName string
	Lineno   int
}

type Goerr interface {
	AddValue(key string, value interface{})

	IsErr() bool
	Message() string
	Route() []RouteNode
	Value(key string) interface{}
}

func newErrRouteNode(file, funcname string, lineno int) RouteNode {
	fn := strings.Split(funcname, ".")
	return RouteNode{
		File:     file,
		FuncName: fn[len(fn)-1],
		Lineno:   lineno,
	}
}

func (e *errTracking) addErrRouteNode(node RouteNode) {
	//	frontNode := e.route
	//	e.cleanErrRouteNode()
	e.route = append(e.route, node)
	//e.route = append(e.route, frontNode...)
}

func (e *errTracking) cleanErrRouteNode() {
	var clear []RouteNode
	e.route = clear
}

func (e *errTracking) setErrMsg(err error) {
	if err != nil {
		e.isErr = true
		e.message = err.Error()
	}
}

func (e *errTracking) setArgs(args ...interface{}) {
	e.args = append(e.args, args...)
}

func Err(err error, args ...interface{}) Goerr {
	var caller int = 1
	goerr := new(errTracking)

	for {
		pc, file, line, ok := runtime.Caller(caller)
		if !ok {
			break
		}

		fn := runtime.FuncForPC(pc)
		if fn.Name() == "runtime.main" {
			break
		} else if fn.Name() == "testing.tRunner" {
			break
		}

		node := newErrRouteNode(file, fn.Name(), line)
		goerr.addErrRouteNode(node)
		caller++
	}

	goerr.setErrMsg(err)
	goerr.setArgs(args...)

	return goerr
}

func (e *errTracking) AddValue(key string, value interface{}) {
	if e.values == nil {
		e.values = make(map[string]interface{})
	}

	if _, ok := e.values[key]; ok {
		fmt.Printf("warn: key %s exist!\n", key)
		return
	}

	e.values[key] = value
}

func (e *errTracking) IsErr() bool {
	return e.isErr
}

func (e *errTracking) Message() string {
	return e.message
}

func (e *errTracking) Route() []RouteNode {
	return e.route
}

func (e *errTracking) Value(key string) interface{} {
	return e.values[key]
}
