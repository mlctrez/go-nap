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

	Page(u *url.URL) Elm

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

type router struct {
	components map[string]ElmFunc
	ops        chan func()
}

func (r *router) E(nodeName string, attr ...M) Elm {
	el := El(nodeName, attr...)
	for _, at := range attr {
		for k, v := range at {
			if k == "href" {
				if _, ok := v.(string); ok {
					el.Listen("click", jsa.FuncOf(func(this jsa.Value, args []jsa.Value) any {
						args[0].PreventDefault()
						u, err := url.Parse(this.Get("href").String())
						if err == nil {
							r.Navigate(u)
						}
						return nil
					}))
					//el.Listen("click", r.NavigateFunc(&url.URL{Path: href}))
				}
			}
		}
	}
	return el
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

func (r *router) ElmFunc(name string, elmFunc ElmFunc) {
	r.components[name] = elmFunc
}

func (r *router) Navigate(u *url.URL) {
	fmt.Println("navigate", u)
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
