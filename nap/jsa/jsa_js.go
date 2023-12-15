package jsa

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"syscall/js"
)

var globalValue = &value{value: js.Global()}
var documentValue = globalValue.Get("document").(*value)
var windowValue = globalValue.Get("window").(*value)
var locationValue = windowValue.Get("location").(*value)

type value struct {
	value js.Value
}

func (v *value) Encode(encoder *xml.Encoder) error {
	panic("not implemented")
}

func (v *value) Set(name string, val any) {
	switch val.(type) {
	case bool:
		v.Call("toggleAttribute", name, val.(bool))
	default:
		v.value.Set(name, ValueOf(val))
	}
}

func (v *value) Get(name string) Value {
	return &value{value: v.value.Get(name)}
}

func (v *value) IsNull() bool {
	return v.value.IsNull()
}

type Func struct {
	fn js.Func
}

func (f Func) Release() {
	f.fn.Release()
}

func ValueOf(arg any) any {
	switch v := arg.(type) {
	case *value:
		return v.value
	case ValueProvider:
		return ValueOf(v.Value())
	case Func:
		return v.fn
	default:
		return js.ValueOf(v)
	}
}

func ValuesOf(args ...any) []any {
	jsArgs := make([]any, len(args))
	for i, arg := range args {
		jsArgs[i] = ValueOf(arg)
	}
	return jsArgs
}

func (v *value) Call(name string, args ...any) Value {
	return &value{value: v.value.Call(name, ValuesOf(args...)...)}
}

func (v *value) String() string {
	return v.value.String()
}

func (v *value) Bool() bool {
	return v.value.Bool()
}

func (v *value) Float() float64 {
	return v.value.Float()
}

func (v *value) PreventDefault() {
	v.Call("preventDefault")
}

func call(name string, args ...any) Value {
	return globalValue.Call(name, args...)
}

func onPopState(cb func()) {
	windowValue.Set("onpopstate", js.FuncOf(func(this js.Value, args []js.Value) any {
		cb()
		return nil
	}))
}

func createElement(name string) Value {
	return documentValue.Call("createElement", name)
}

func createElementNS(ns, name string) Value {
	fmt.Println("createElementNS")
	return documentValue.Call("createElementNS", ns, name)
}

func createTextNode(data string) Value {
	return documentValue.Call("createTextNode", data)
}

func body() Value {
	return documentValue.Get("body")
}

func global() Value {
	return globalValue
}

func currentUrl() *url.URL {
	href := js.Global().Get("window").Get("location").Get("href").String()
	if u, err := url.Parse(href); err != nil {
		panic(err)
	} else {
		return u
	}
}

func elById(id string) Value {
	return documentValue.Call("getElementById", id)
}

func addHistory(href string) {
	if href != locationValue.Get("href").String() {
		windowValue.Get("history").Call("pushState", nil, "", href)
	}
}

func funcOf(callback func(this Value, args []Value) any) Func {
	fn := js.FuncOf(func(this js.Value, args []js.Value) any {
		jsaArgs := make([]Value, len(args))
		for i, arg := range args {
			jsaArgs[i] = &value{value: arg}
		}
		return callback(&value{value: this}, jsaArgs)
	})
	return Func{fn: fn}
}

func localStorage() LocalStorageApi {
	return &localStorageImpl{value: globalValue.Get("localStorage").(*value)}
}

type localStorageImpl struct {
	value *value
}

func (l *localStorageImpl) Set(name string, value any) {
	l.value.Call("setItem", name, value)
}

func (l *localStorageImpl) Get(name string) Value {
	return l.value.Call("getItem", name)
}

func (l *localStorageImpl) GetS(name string) string {
	return l.Get(name).String()
}

func (l *localStorageImpl) Remove(name string) {
	l.value.Call("removeItem", name)
}
