package components

import "github.com/mlctrez/go-nap/nap"

func CompoOverride(r nap.Router) {
	r.Override(CompoDropdownMenu, func(r nap.Router) nap.Elm {
		menu := r.ElmOrig(CompoDropdownMenu)
		for _, option := range []string{"a", "b", "c"} {
			menu.Append(
				r.Elm(CompoDropdownItem).
					Find("a").Set("href", "/").Append(nap.Text(option)))
		}
		return menu
	})
}
