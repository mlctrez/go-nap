package demo

import (
	"github.com/mlctrez/go-nap/demo/components"
	"github.com/mlctrez/go-nap/demo/nav"
	"github.com/mlctrez/go-nap/demo/web"
	"github.com/mlctrez/go-nap/nap"
)

func Routes(r nap.Router) {
	r.With(
		web.Base,
		nav.Navbar,
		components.Compo,
	)

	//r.ElmFunc("_page/", func(r nap.Router) nap.Elm {
	//	return r.Elm(web.BaseHtml).Append(r.Elm("_body/"))
	//})
	//r.ElmFunc("_body/", func(r nap.Router) nap.Elm {
	//	return r.E("body").Append(
	//		r.Elm(nav.NavbarNav),
	//		r.Elm(components.CompoDropdown),
	//	)
	//})
}
