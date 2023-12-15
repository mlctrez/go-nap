package compo

import "github.com/mlctrez/go-nap/nap"

func DropdownOverride(r nap.Router) {
	r.Override(CDropdownMenu, func(r nap.Router) nap.Elm {
		menu := r.ElmOrig(CDropdownMenu)
		for _, option := range []string{"a", "b", "c"} {
			menu.Append(r.NavLink(r.Elm(CDropdownDropdownItem), "/", option))
		}
		return menu
	})
}
