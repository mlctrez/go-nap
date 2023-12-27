package web

import (
	"github.com/mlctrez/go-nap/demo/compo"
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/go-nap/nap/jsa"
)

func BaseOverride(r nap.Router) {

	// how to handle logged in vs. not logged in?
	// some pages should redirect to the login page if the user is not logged in
	// other pages can display normally without being logged in
	// should there be an attribute on the page that designates that a local storage value should
	//   be present in order for the page to be rendered

	pageAndBody(r, "/", func(r nap.Router) nap.Elm {
		user := jsa.LocalStorage().Get("user")
		if !user.IsNull() {
			return r.E("body").Append(r.Elm(compo.ENavbarNav), r.Elm(compo.EDropdownMain))
		}
		return r.Elm(compo.ESignInBody)
	})
	pageAndBody(r, "/other", func(r nap.Router) nap.Elm { return r.E("body").Append(r.Elm(compo.ENavbarNav)) })
}

func pageAndBody(r nap.Router, u string, elmFunc nap.ElmFunc) {
	r.ElmFunc("_page"+u, func(r nap.Router) nap.Elm { return r.Elm(EBaseHtml).Append(r.Elm("_body" + u)) })
	r.ElmFunc("_body"+u, elmFunc)
}
