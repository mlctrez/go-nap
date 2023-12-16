// This code is auto generated. DO NOT EDIT.

package compo

import nap "github.com/mlctrez/go-nap/nap"

const ENavbarNav = "navbar/nav"
const ENavbarItems = "navbar/items"
const ENavbarActiveNavItem = "navbar/activeNavItem"
const ENavbarInactiveNavItem = "navbar/inactiveNavItem"
const ENavbarDisabledNavItem = "navbar/disabledNavItem"
const ENavbarDropdownNavItem = "navbar/dropdownNavItem"
const ENavbarDropdownMenu = "navbar/dropdownMenu"
const ENavbarDropdownItem = "navbar/dropdownItem"
const ENavbarDropdownDivider = "navbar/dropdownDivider"
const ENavbarSearchForm = "navbar/searchForm"

func Navbar(r nap.Router) {
	r.ElmFunc(ENavbarNav, NavbarNav)
	r.ElmFunc(ENavbarItems, NavbarItems)
	r.ElmFunc(ENavbarActiveNavItem, NavbarActiveNavItem)
	r.ElmFunc(ENavbarInactiveNavItem, NavbarInactiveNavItem)
	r.ElmFunc(ENavbarDisabledNavItem, NavbarDisabledNavItem)
	r.ElmFunc(ENavbarDropdownNavItem, NavbarDropdownNavItem)
	r.ElmFunc(ENavbarDropdownMenu, NavbarDropdownMenu)
	r.ElmFunc(ENavbarDropdownItem, NavbarDropdownItem)
	r.ElmFunc(ENavbarDropdownDivider, NavbarDropdownDivider)
	r.ElmFunc(ENavbarSearchForm, NavbarSearchForm)
	NavbarOverride(r)
}

func NavbarNav(r nap.Router) nap.Elm {
	return r.E("nav").Set("class", "navbar navbar-expand-lg bg-body-tertiary").
		Append(r.E("div").Set("class", "container-fluid").
			Append(
				r.E("a").Set("class", "navbar-brand").Set("href", "/").
					Append(nap.Text("Navbar")),
				r.E("button").
					Set("class", "navbar-toggler").
					Set("type", "button").
					Set("data-bs-toggle", "collapse").
					Set("data-bs-target", "#navbarSupportedContent").
					Set("aria-controls", "navbarSupportedContent").
					Set("aria-expanded", "false").
					Set("aria-label", "Toggle navigation").
					Append(r.E("span").Set("class", "navbar-toggler-icon")),
				r.E("div").Set("class", "collapse navbar-collapse").Set("id", "navbarSupportedContent").
					Append(
						r.Elm(ENavbarItems),
						r.Elm(ENavbarSearchForm))))
}

func NavbarItems(r nap.Router) nap.Elm {
	return r.E("ul").Set("class", "navbar-nav me-auto mb-2 mb-lg-0")
}

func NavbarActiveNavItem(r nap.Router) nap.Elm {
	return r.E("li").Set("class", "nav-item").
		Append(r.E("a").Set("class", "nav-link active").Set("aria-current", "page"))
}

func NavbarInactiveNavItem(r nap.Router) nap.Elm {
	return r.E("li").Set("class", "nav-item").
		Append(r.E("a").Set("class", "nav-link"))
}

func NavbarDisabledNavItem(r nap.Router) nap.Elm {
	return r.E("li").Set("class", "nav-item").
		Append(r.E("a").Set("class", "nav-link disabled").Set("aria-disabled", "true"))
}

func NavbarDropdownNavItem(r nap.Router) nap.Elm {
	return r.E("li").Set("class", "nav-item dropdown").
		Append(
			r.E("a").
				Set("class", "nav-link dropdown-toggle").
				Set("role", "button").
				Set("data-bs-toggle", "dropdown").
				Set("aria-expanded", "false").
				Append(nap.Text("Dropdown")),
			r.Elm(ENavbarDropdownMenu))
}

func NavbarDropdownMenu(r nap.Router) nap.Elm {
	return r.E("ul").Set("class", "dropdown-menu")
}

func NavbarDropdownItem(r nap.Router) nap.Elm {
	return r.E("li").
		Append(r.E("a").Set("class", "dropdown-item").Set("href", "/").
			Append(nap.Text("Action")))
}

func NavbarDropdownDivider(r nap.Router) nap.Elm {
	return r.E("li").
		Append(r.E("hr").Set("class", "dropdown-divider"))
}

func NavbarSearchForm(r nap.Router) nap.Elm {
	return r.E("form").
		Set("id", "searchForm").
		Set("class", "d-flex").
		Set("role", "search").
		Append(
			r.E("input").
				Set("name", "searchFormInput").
				Set("class", "form-control me-2").
				Set("type", "search").
				Set("placeholder", "Search").
				Set("aria-label", "Search"),
			r.E("button").Set("class", "btn btn-outline-success").Set("type", "submit").
				Append(nap.Text("Search")))
}
