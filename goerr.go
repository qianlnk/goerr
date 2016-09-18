package goerr

import (
	"fmt"
	"runtime"
	"strings"
)

type errTracking struct {
	route   []RouteNode
	message string
	values  map[string]interface{}
}

type RouteNode struct {
	File     string
	FuncName string
	Lineno   int
}

type Goerr interface {
	AddValue(key string, value interface{})

	Message() string
	Route() []RouteNode
	Value(key string) interface{}

	Stdout() //show fmt message for debug
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
	e.message = err.Error()
}

func Err(err error) Goerr {
	if err == nil {
		return nil
	}

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

func (e *errTracking) Message() string {
	return e.message
}

func (e *errTracking) Route() []RouteNode {
	return e.route
}

func (e *errTracking) Value(key string) interface{} {
	return e.values[key]
}

func (e *errTracking) Stdout() {
	fmt.Println("Error:", e.message)
	fmt.Println("Error Route:")
	for _, r := range e.route {
		fmt.Println(r.File, r.Lineno, r.FuncName)
	}

	for k, v := range e.values {
		fmt.Println(k, v)
	}
}
