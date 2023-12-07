package nav

import nap "github.com/mlctrez/go-nap/nap"

const NavbarNav = "navbar/nav"
const NavbarItems = "navbar/items"
const NavbarActiveNavItem = "navbar/activeNavItem"
const NavbarInactiveNavItem = "navbar/inactiveNavItem"
const NavbarDisabledNavItem = "navbar/disabledNavItem"
const NavbarDropdownNavItem = "navbar/dropdownNavItem"
const NavbarDropdownMenu = "navbar/dropdownMenu"
const NavbarDropdownItem = "navbar/dropdownItem"
const NavbarDropdownDivider = "navbar/dropdownDivider"
const NavbarSearchForm = "navbar/searchForm"

func Navbar(r nap.Router) {
	r.ElmFunc(NavbarNav, Nav)
	r.ElmFunc(NavbarItems, Items)
	r.ElmFunc(NavbarActiveNavItem, ActiveNavItem)
	r.ElmFunc(NavbarInactiveNavItem, InactiveNavItem)
	r.ElmFunc(NavbarDisabledNavItem, DisabledNavItem)
	r.ElmFunc(NavbarDropdownNavItem, DropdownNavItem)
	r.ElmFunc(NavbarDropdownMenu, DropdownMenu)
	r.ElmFunc(NavbarDropdownItem, DropdownItem)
	r.ElmFunc(NavbarDropdownDivider, DropdownDivider)
	r.ElmFunc(NavbarSearchForm, SearchForm)
	NavbarOverride(r)
}

func Nav(r nap.Router) nap.Elm {
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
						r.Elm(NavbarItems),
						r.Elm(NavbarSearchForm))))
}

func Items(r nap.Router) nap.Elm {
	return r.E("ul").Set("class", "navbar-nav me-auto mb-2 mb-lg-0")
}

func ActiveNavItem(r nap.Router) nap.Elm {
	return r.E("li").Set("class", "nav-item").
		Append(r.E("a").Set("class", "nav-link active").Set("aria-current", "page"))
}

func InactiveNavItem(r nap.Router) nap.Elm {
	return r.E("li").Set("class", "nav-item").
		Append(r.E("a").Set("class", "nav-link"))
}

func DisabledNavItem(r nap.Router) nap.Elm {
	return r.E("li").Set("class", "nav-item").
		Append(r.E("a").Set("class", "nav-link disabled").Set("aria-disabled", "true"))
}

func DropdownNavItem(r nap.Router) nap.Elm {
	return r.E("li").Set("class", "nav-item dropdown").
		Append(
			r.E("a").
				Set("class", "nav-link dropdown-toggle").
				Set("role", "button").
				Set("data-bs-toggle", "dropdown").
				Set("aria-expanded", "false").
				Append(nap.Text("Dropdown")),
			r.Elm(NavbarDropdownMenu))
}

func DropdownMenu(r nap.Router) nap.Elm {
	return r.E("ul").Set("class", "dropdown-menu")
}

func DropdownItem(r nap.Router) nap.Elm {
	return r.E("li").
		Append(r.E("a").Set("class", "dropdown-item").Set("href", "/").
			Append(nap.Text("Action")))
}

func DropdownDivider(r nap.Router) nap.Elm {
	return r.E("li").
		Append(r.E("hr").Set("class", "dropdown-divider"))
}

func SearchForm(r nap.Router) nap.Elm {
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
