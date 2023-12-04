package web

import (
	"github.com/mlctrez/go-nap/demo/nav"
	"github.com/mlctrez/go-nap/nap"
)

func Routes(r nap.Router) {
	r.With(Base, nav.Navbar)
	r.ElmFunc("_page/", func(r nap.Router) nap.Elm { return r.Elm(BaseHtml).Append(r.Elm(nav.NavbarNav)) })
	r.ElmFunc("_body/", func(r nap.Router) nap.Elm { return r.E("body").Append(r.Elm(nav.NavbarNav)) })
}
