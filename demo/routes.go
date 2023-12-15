package demo

import (
	"github.com/mlctrez/go-nap/demo/compo"
	"github.com/mlctrez/go-nap/demo/web"
	"github.com/mlctrez/go-nap/nap"
)

func Routes(r nap.Router) {
	r.With(
		web.Base,
		compo.Navbar,
		compo.Dropdown,
		compo.SignIn,
	)
}
