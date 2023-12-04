//go:build !js

package jsa

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

type value struct {
	nodeName string
	attr     map[string]any
	children []*value
}

func (v *value) Encode(enc *xml.Encoder) error {

	if v.nodeName == "#text" {
		return enc.EncodeToken(xml.CharData(v.attr["data"].(string)))
	}

	name := xml.Name{Local: v.nodeName}
	var attrs []xml.Attr
	for k, v := range v.attr {
		attrs = append(attrs, xml.Attr{Name: xml.Name{Local: k}, Value: fmt.Sprintf("%v", v)})
	}
	err := enc.EncodeToken(xml.StartElement{Name: name, Attr: attrs})
	if err != nil {
		return err
	}

	for _, child := range v.children {
		err = child.Encode(enc)
		if err != nil {
			return err
		}
	}

	return enc.EncodeToken(xml.EndElement{Name: name})
}

func call(name string, args ...any) Value {
	return &value{}
}

func onPopState(cb func()) {
}

type Func struct {
}

func (f Func) Release() {
}

func createElement(name string) Value {
	return &value{nodeName: name}
}

func createElementNS(ns, name string) Value {
	return &value{nodeName: name}
}

func createTextNode(data string) Value {
	v := &value{nodeName: "#text"}
	if v.attr == nil {
		v.attr = make(map[string]any)
	}
	v.attr["data"] = data
	return v
}

func funcOf(input func(this Value, args []Value) any) Func {
	return Func{}
}

func body() Value {
	return &value{}
}

func global() Value {
	return &value{}
}

func currentUrl() *url.URL {
	return &url.URL{}
}

func addHistory(href string) {}

func elById(id string) Value {
	return &value{}
}

func (v *value) Set(name string, val any) {
	if v.attr == nil {
		v.attr = make(map[string]any)
	}
	v.attr[name] = val
}

func (v *value) Get(name string) Value {
	return &value{}
}

func (v *value) IsNull() bool {
	return true
}

func (v *value) String() string {
	return fmt.Sprintf("node %q", v.nodeName)
}

func (v *value) Call(name string, args ...any) Value {
	switch name {
	case "replaceWith":
		switch t := args[0].(type) {
		case *value:
			*v = *t
		}
	case "appendChild":
		for _, arg := range args {
			v.children = append(v.children, arg.(*value))
		}
	case "setAttribute":
		v.Set(args[0].(string), args[1])
	}
	return v
}

func (v *value) PreventDefault() {
}

func localStorage() LocalStorageApi {
	return &localStorageImpl{}
}

type localStorageImpl struct {
}

func (l *localStorageImpl) Set(name string, value any) {
}

func (l *localStorageImpl) Get(name string) Value {
	return &value{}
}
func (l *localStorageImpl) GetS(name string) string {
	return l.Get(name).String()
}

func (l *localStorageImpl) Remove(name string) {
}
