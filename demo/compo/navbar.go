package compo

import nap "github.com/mlctrez/go-nap/nap"

const CNavbarNav = "navbar/nav"
const CNavbarItems = "navbar/items"
const CNavbarActiveNavItem = "navbar/activeNavItem"
const CNavbarInactiveNavItem = "navbar/inactiveNavItem"
const CNavbarDisabledNavItem = "navbar/disabledNavItem"
const CNavbarDropdownNavItem = "navbar/dropdownNavItem"
const CNavbarDropdownMenu = "navbar/dropdownMenu"
const CNavbarDropdownItem = "navbar/dropdownItem"
const CNavbarDropdownDivider = "navbar/dropdownDivider"
const CNavbarSearchForm = "navbar/searchForm"

func Navbar(r nap.Router) {
	r.ElmFunc(CNavbarNav, NavbarNav)
	r.ElmFunc(CNavbarItems, NavbarItems)
	r.ElmFunc(CNavbarActiveNavItem, NavbarActiveNavItem)
	r.ElmFunc(CNavbarInactiveNavItem, NavbarInactiveNavItem)
	r.ElmFunc(CNavbarDisabledNavItem, NavbarDisabledNavItem)
	r.ElmFunc(CNavbarDropdownNavItem, NavbarDropdownNavItem)
	r.ElmFunc(CNavbarDropdownMenu, NavbarDropdownMenu)
	r.ElmFunc(CNavbarDropdownItem, NavbarDropdownItem)
	r.ElmFunc(CNavbarDropdownDivider, NavbarDropdownDivider)
	r.ElmFunc(CNavbarSearchForm, NavbarSearchForm)
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
						r.Elm(CNavbarItems),
						r.Elm(CNavbarSearchForm))))
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
			r.Elm(CNavbarDropdownMenu))
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
