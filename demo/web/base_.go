package web

import (
	"github.com/mlctrez/go-nap/demo/compo"
	"github.com/mlctrez/go-nap/nap"
)

func BaseOverride(r nap.Router) {
	pageAndBody(r, "/", func(r nap.Router) nap.Elm {
		return r.E("body").Set("class", "d-flex align-items-center py-4 bg-body-tertiary").Append(
			r.Elm(compo.CSignInMain),
			//r.Elm(nav.NavbarNav),
			//r.Elm(components.CompoDropdown),
		)
	})
	pageAndBody(r, "/other", func(r nap.Router) nap.Elm {
		return r.E("body").Append(
			r.Elm(compo.CNavbarNav),
		)
	})
}

func pageAndBody(r nap.Router, u string, elmFunc nap.ElmFunc) {
	r.ElmFunc("_page"+u, func(r nap.Router) nap.Elm { return r.Elm(CBaseHtml).Append(r.Elm("_body" + u)) })
	r.ElmFunc("_body"+u, elmFunc)
}
