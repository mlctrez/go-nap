package jsa

import (
	"encoding/xml"
	"net/url"
)

type ValueProvider interface {
	Value() Value
}

func Call(name string, args ...any) Value {
	return call(name, args...)
}

func OnPopState(cb func()) {
	onPopState(cb)
}

func CreateElement(name string) Value {
	return createElement(name)
}

func CreateElementNS(ns, name string) Value {
	return createElementNS(ns, name)
}

func CreateTextNode(data string) Value {
	return createTextNode(data)
}

func Body() Value {
	return body()
}

func Global() Value {
	return global()
}

func CurrentURL() *url.URL {
	return currentUrl()
}

func ElById(id string) Value {
	return elById(id)
}

func ValueById(id string) string {
	return elById(id).Get("value").String()
}

func AddHistory(href string) {
	addHistory(href)
}

type Value interface {
	Call(name string, args ...any) Value
	Get(name string) Value
	Set(name string, val any)
	String() string
	Encode(encoder *xml.Encoder) error
	PreventDefault()
	IsNull() bool
}

func FuncOf(input func(this Value, args []Value) any) Func {
	return funcOf(input)
}

func Log(args ...any) {
	Call("napConsoleLog", args...)
}

type LocalStorageApi interface {
	Set(name string, value any)
	Get(name string) Value
	GetS(name string) string
	Remove(name string)
}

func LocalStorage() LocalStorageApi {
	return localStorage()
}
