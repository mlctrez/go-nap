// This code is auto generated. DO NOT EDIT.

package compo

import nap "github.com/mlctrez/go-nap/nap"

const EDropdownMain = "dropdown/main"
const EDropdownMenu = "dropdown/menu"
const EDropdownDropdownItem = "dropdown/dropdownItem"

func Dropdown(r nap.Router) {
	r.ElmFunc(EDropdownMain, DropdownMain)
	r.ElmFunc(EDropdownMenu, DropdownMenu)
	r.ElmFunc(EDropdownDropdownItem, DropdownDropdownItem)
	DropdownOverride(r)
}

func DropdownMain(r nap.Router) nap.Elm {
	return r.E("div").Set("class", "dropdown").
		Append(
			r.E("button").
				Set("class", "btn btn-primary dropdown-toggle").
				Set("type", "button").
				Set("data-bs-toggle", "dropdown").
				Set("aria-expanded", "false").
				Append(nap.Text("Dropdown")),
			r.Elm(EDropdownMenu))
}

func DropdownMenu(r nap.Router) nap.Elm {
	return r.E("ul").Set("class", "dropdown-menu")
}

func DropdownDropdownItem(r nap.Router) nap.Elm {
	return r.E("li").
		Append(r.E("a").Set("class", "dropdown-item").Set("href", "#"))
}
