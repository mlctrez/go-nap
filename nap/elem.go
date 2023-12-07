package nap

import (
	"encoding/xml"
	"fmt"
	"github.com/mlctrez/go-nap/nap/enc"
	"github.com/mlctrez/go-nap/nap/jsa"
	"runtime"
	"strings"
)

type Elm interface {
	NodeName() string

	Set(name string, val any) Elm
	Get(name string) jsa.Value
	Value() jsa.Value

	Listen(name string, cb jsa.Func) Elm

	Append(el ...Elm) Elm
	Text(data string) Elm

	Find(name ...string) Elm
	FindPath(path string) Elm
	First(name string) Elm
	All(name string) []Elm
	Children() []Elm
	ReplaceChild(old, new Elm)

	Encode(encoder *xml.Encoder) error
	String() string
}

type M map[string]any

type ElmFunc func(r Router) Elm

type elem struct {
	nodeName  string
	node      jsa.Value
	children  []Elm
	events    []jsa.Func
	fnRelease jsa.Func
}

func elFinalizer(e *elem) {
	fmt.Println("finalizing", e.nodeName)
}

func El(nodeName string, attr ...M) Elm {
	el := &elem{nodeName: nodeName, node: jsa.CreateElement(nodeName)}
	runtime.SetFinalizer(el, elFinalizer)
	return el.SetAttr(attr...)
}

func ElNS(ns, nodeName string, attr ...M) Elm {
	el := &elem{nodeName: nodeName, node: jsa.CreateElementNS(ns, nodeName)}
	runtime.SetFinalizer(el, elFinalizer)
	return el.SetAttr(attr...)
}

func (e *elem) NodeName() string {
	return e.nodeName
}

func (e *elem) Encode(encoder *xml.Encoder) error {
	return e.node.Encode(encoder)
}

func (e *elem) String() string {
	encoder := enc.New().Indent("  ")
	if err := encoder.Encode(e); err != nil {
		return err.Error()
	}
	return encoder.Content()
}

func (e *elem) ReplaceChild(old, new Elm) {
	children := e.Children()
	for i, elm := range children {
		if elm == old {
			children[i] = new
			old.Value().Call("replaceWith", new.Value())
		}
	}
}

func (e *elem) SetAttr(attr ...M) Elm {
	for _, m := range attr {
		for k, v := range m {
			e.Set(k, v)
		}
	}
	return e
}

func (e *elem) Set(name string, val any) Elm {
	if name == "value" {
		e.node.Set(name, val)
		return e
	}
	switch val.(type) {
	case bool:
		e.node.Call("toggleAttribute", name, val.(bool))
	default:
		e.node.Call("setAttribute", name, val)
	}
	return e
}

func (e *elem) Get(name string) jsa.Value {
	return e.node.Get(name)
}

func (e *elem) Append(el ...Elm) Elm {
	for _, child := range el {
		e.node.Call("appendChild", child.Value())
		e.children = append(e.children, child)
	}
	return e
}

func (e *elem) Children() []Elm {
	return e.children
}

func (e *elem) FindPath(path string) Elm {
	return e.Find(strings.Split(path, "/")...)
}

func (e *elem) Find(name ...string) Elm {
	if len(name) == 0 {
		panic("cannot find without at least one tag name")
	}
	var next Elm
	if len(name) > 1 {
		if next = e.First(name[0]); next != nil {
			return next.Find(name[1:]...)
		}
	}
	return e.First(name[0])
}

func (e *elem) First(name string) Elm {
	if e.NodeName() == name {
		return e
	}
	for _, child := range e.children {
		if child.NodeName() == name {
			return child
		}
		if ch := child.First(name); ch != nil {
			return ch
		}
	}
	return nil
}

func (e *elem) All(name string) []Elm {
	var result []Elm
	if e.NodeName() == name {
		result = append(result, e)
	}
	for _, child := range e.children {
		result = append(result, child.All(name)...)
	}
	return result
}

func (e *elem) Prepend(el ...Elm) Elm {
	for _, child := range el {
		e.node.Call("prepend", child.Value())
		e.children = append([]Elm{child}, e.children...)
	}
	return e
}

func (e *elem) Value() jsa.Value {
	return e.node
}

func (e *elem) Text(data string) Elm {
	e.Append(Text(data))
	return e
}

func Text(data string) Elm {
	return &elem{nodeName: "#text", node: jsa.CreateTextNode(data)}
}

func (e *elem) Listen(name string, fn jsa.Func) Elm {
	e.node.Call("addEventListener", name, e.addRelease(fn))
	return e
}

func (e *elem) addRelease(fn jsa.Func) jsa.Func {
	if len(e.events) == 0 {
		e.fnRelease = jsa.FuncOf(func(_ jsa.Value, _ []jsa.Value) any {
			for _, ev := range e.events {
				ev.Release()
			}
			e.fnRelease.Release()
			return nil
		})
		e.node.Set("napRelease", e.fnRelease)
	}
	e.events = append(e.events, fn)
	return fn
}
