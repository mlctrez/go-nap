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
}
