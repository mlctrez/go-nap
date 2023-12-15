package compo

import (
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/go-nap/nap/jsa"
	"net/url"
)

func SignInOverride(r nap.Router) {
	r.Override(CSignInMain, func(r nap.Router) nap.Elm {
		m := r.ElmOrig(CSignInMain)
		m.First("form").Listen("submit", jsa.FuncOf(func(this jsa.Value, args []jsa.Value) any {
			args[0].PreventDefault()
			if m.First("input").Value().Get("value").String() != "" {
				r.Navigate(&url.URL{Path: "/other"})
			}
			return nil
		}))
		return m
	})
}
