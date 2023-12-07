package web

import (
	"github.com/mlctrez/go-nap/demo/components"
	"github.com/mlctrez/go-nap/demo/nav"
	"github.com/mlctrez/go-nap/nap"
)

func BaseOverride(r nap.Router) {

	// TODO: move these into html declarations

	r.ElmFunc("_page/", func(r nap.Router) nap.Elm {
		return r.Elm(BaseHtml).Append(r.Elm("_body/"))
	})
	r.ElmFunc("_body/", func(r nap.Router) nap.Elm {
		return r.E("body").Append(
			r.Elm(nav.NavbarNav),
			r.Elm(components.CompoDropdown),
		)
	})

}
