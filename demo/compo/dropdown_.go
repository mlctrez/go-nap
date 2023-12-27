package compo

import "github.com/mlctrez/go-nap/nap"

func DropdownOverride(r nap.Router) {
	r.Override(EDropdownMenu, func(r nap.Router) nap.Elm {
		menu := r.ElmOrig(EDropdownMenu)
		for _, option := range []string{"option a", "option b", "option c", "option d", "option e"} {
			menu.Append(r.NavLink(r.Elm(EDropdownDropdownItem), "/", option))
		}
		return menu
	})
}
