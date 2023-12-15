package nap

import (
	_ "embed"
	"fmt"
	"github.com/mlctrez/go-nap/nap/jsa"
	"net/url"
	"strings"
)

type Router interface {
	Navigate(u *url.URL)
	NavigateFunc(u *url.URL) jsa.Func

	QueueOp(f func())
	Ops() chan func()

	With(rf ...RegFunc) Router

	ElmFunc(name string, elmFunc ElmFunc)
	Elm(name string) Elm
	ElmOrig(name string) Elm

	Page(u *url.URL) Elm

	Override(name string, elmFunc ElmFunc)
	NavLink(el Elm, u, text string) Elm

	E(nodeName string, attr ...M) Elm
}

type RegFunc func(r Router)

func NewRouter(rf ...RegFunc) Router {
	r := &router{
		components: make(map[string]ElmFunc),
		ops:        make(chan func(), 20),
	}
	r.With(rf...)
	return r
}

var _ Router = (*router)(nil)

type router struct {
	components map[string]ElmFunc
	ops        chan func()
}

func (r *router) NavLink(el Elm, u, text string) Elm {
	anchor := el.First("a")
	if len(anchor.Children()) == 0 {
		anchor.Append(Text(text))
	} else {
		anchor.ReplaceChild(anchor.Children()[0], Text(text))
	}
	anchor.Set("href", u)
	anchor.Listen("click", jsa.FuncOf(func(this jsa.Value, args []jsa.Value) any {
		args[0].PreventDefault()
		r.QueueOp(func() { r.Navigate(&url.URL{Path: u}) })
		return nil
	}))
	return el
}

func (r *router) E(nodeName string, attr ...M) Elm {
	return El(nodeName, attr...)
}

func (r *router) Page(u *url.URL) Elm {
	return r.Elm("_page" + u.Path)
}

func (r *router) Body(u *url.URL) Elm {
	return r.Elm("_body" + u.Path)
}

func (r *router) With(rf ...RegFunc) Router {
	for _, regFunc := range rf {
		regFunc(r)
	}
	return r
}

func (r *router) Elm(name string) Elm {
	if ef, ok := r.components[name]; ok {
		return ef(r)
	}
	if strings.HasPrefix(name, "_body") {
		return El("body", M{"style": "color:red"}).
			Text(fmt.Sprintf("element %q not found", name))
	}
	return El("div", M{"style": "color:red"}).
		Text(fmt.Sprintf("element %q not found", name))
}

func (r *router) ElmOrig(name string) Elm {
	if ef, ok := r.components[name+"Orig"]; ok {
		return ef(r)
	}
	if strings.HasPrefix(name, "_body") {
		return El("body", M{"style": "color:red"}).
			Text(fmt.Sprintf("element %q not found", name))
	}
	return El("div", M{"style": "color:red"}).
		Text(fmt.Sprintf("element %q not found", name))
}

func (r *router) ElmFunc(name string, elmFunc ElmFunc) {
	r.components[name] = elmFunc
	r.components[name+"Orig"] = elmFunc
}

func (r *router) Override(name string, override ElmFunc) {
	if _, ok := r.components[name]; !ok {
		panic(fmt.Sprintf("component %q does not exist", name))
	} else {
		r.components[name] = override
	}
}

func (r *router) Navigate(u *url.URL) {
	//fmt.Println("navigate", u)
	jsa.AddHistory(u.String())
	jsa.Body().Call("replaceWith", r.Body(u))
}

func (r *router) NavigateFunc(u *url.URL) jsa.Func {
	return jsa.FuncOf(func(this jsa.Value, args []jsa.Value) any {
		args[0].PreventDefault()
		r.QueueOp(func() { r.Navigate(u) })
		return nil
	})
}

func (r *router) QueueOp(f func()) {
	r.ops <- f
}

func (r *router) Ops() chan func() {
	return r.ops
}
