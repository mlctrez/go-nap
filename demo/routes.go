package demo

import (
	"github.com/mlctrez/go-nap/demo/nav"
	"github.com/mlctrez/go-nap/demo/web"
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/go-nap/nap/jsa"
)

func Routes(r nap.Router) {
	r.With(web.Base, nav.Navbar)
	r.ElmFunc("_page/", func(r nap.Router) nap.Elm {
		return r.Elm(web.BaseHtml).Append(r.Elm("_body/"))
	})
	r.ElmFunc("_body/", func(r nap.Router) nap.Elm {
		elm := r.Elm(nav.NavbarNav)
		searchForm := r.Elm(nav.NavbarSearchForm)
		searchForm.Listen("submit", jsa.FuncOf(func(this jsa.Value, args []jsa.Value) any {
			args[0].PreventDefault()
			jsa.Log(searchForm.First("input").Get("value"))
			return nil
		}))
		elm.FindPath("nav/div/div").Append(searchForm)
		return r.E("body").Append(elm)
	})
}
