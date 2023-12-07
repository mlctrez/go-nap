package web

import (
	"github.com/mlctrez/go-nap/demo/components"
	"github.com/mlctrez/go-nap/demo/nav"
	"github.com/mlctrez/go-nap/nap"
)

func BaseOverride(r nap.Router) {
	pageAndBody(r, "/", func(r nap.Router) nap.Elm {
		return r.E("body").Append(
			r.Elm(nav.NavbarNav),
			r.Elm(components.CompoDropdown),
		)
	})
	pageAndBody(r, "/other", func(r nap.Router) nap.Elm {
		return r.E("body").Append(
			r.Elm(nav.NavbarNav),
		)
	})
}

func pageAndBody(r nap.Router, u string, elmFunc nap.ElmFunc) {
	r.ElmFunc("_page"+u, func(r nap.Router) nap.Elm { return r.Elm(BaseHtml).Append(r.Elm("_body" + u)) })
	r.ElmFunc("_body"+u, elmFunc)
}
