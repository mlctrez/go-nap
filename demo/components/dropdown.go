package components

import nap "github.com/mlctrez/go-nap/nap"

const CompoDropdown = "compo/dropdown"
const CompoDropdownMenu = "compo/dropdownMenu"
const CompoDropdownItem = "compo/dropdownItem"

func Compo(r nap.Router) {
	r.ElmFunc(CompoDropdown, Dropdown)
	r.ElmFunc(CompoDropdownMenu, DropdownMenu)
	r.ElmFunc(CompoDropdownItem, DropdownItem)
	CompoOverride(r)
}

func Dropdown(r nap.Router) nap.Elm {
	return r.E("div").Set("class", "dropdown").
		Append(
			r.E("button").
				Set("class", "btn btn-secondary dropdown-toggle").
				Set("type", "button").
				Set("data-bs-toggle", "dropdown").
				Set("aria-expanded", "false").
				Append(nap.Text("Dropdown button")),
			r.Elm(CompoDropdownMenu))
}

func DropdownMenu(r nap.Router) nap.Elm {
	return r.E("ul").Set("class", "dropdown-menu")
}

func DropdownItem(r nap.Router) nap.Elm {
	return r.E("li").
		Append(r.E("a").Set("class", "dropdown-item").Set("href", "#"))
}
