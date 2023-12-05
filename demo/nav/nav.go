package nav

import nap "github.com/mlctrez/go-nap/nap"

const NavbarNav = "navbar/nav"
const NavbarItems = "navbar/items"
const NavbarActiveNavItem = "navbar/activeNavItem"
const NavbarInactiveNavItem = "navbar/inactiveNavItem"
const NavbarDropdownNavItem = "navbar/dropdownNavItem"
const NavbarDropdownMenu = "navbar/dropdownMenu"
const NavbarDisabledNavItem = "navbar/disabledNavItem"
const NavbarSearchForm = "navbar/searchForm"

func Navbar(r nap.Router) {
	r.ElmFunc(NavbarNav, Nav)
	r.ElmFunc(NavbarItems, Items)
	r.ElmFunc(NavbarActiveNavItem, ActiveNavItem)
	r.ElmFunc(NavbarInactiveNavItem, InactiveNavItem)
	r.ElmFunc(NavbarDropdownNavItem, DropdownNavItem)
	r.ElmFunc(NavbarDropdownMenu, DropdownMenu)
	r.ElmFunc(NavbarDisabledNavItem, DisabledNavItem)
	r.ElmFunc(NavbarSearchForm, SearchForm)
}

func Nav(r nap.Router) nap.Elm {
	return r.
		E("nav", nap.M{"class": "navbar navbar-expand-lg bg-body-tertiary"}).
		Append(r.
			E("div", nap.M{"class": "container-fluid"}).
			Append(r.
				E("a", nap.M{"class": "navbar-brand", "href": "/"}).
				Append(nap.Text("Navbar")), r.
				E("button", nap.M{
					"class":          "navbar-toggler",
					"type":           "button",
					"data-bs-toggle": "collapse",
					"data-bs-target": "#navbarSupportedContent",
					"aria-controls":  "navbarSupportedContent",
					"aria-expanded":  "false",
					"aria-label":     "Toggle navigation",
				}).
				Append(r.
					E("span", nap.M{"class": "navbar-toggler-icon"})), r.
				E("div", nap.M{"class": "collapse navbar-collapse", "id": "navbarSupportedContent"}).
				Append(r.
					Elm(NavbarItems))))
}

func Items(r nap.Router) nap.Elm {
	return r.
		E("ul", nap.M{"class": "navbar-nav me-auto mb-2 mb-lg-0"})
}

func ActiveNavItem(r nap.Router) nap.Elm {
	return r.
		E("li", nap.M{"class": "nav-item"}).
		Append(r.
			E("a", nap.M{"class": "nav-link active", "aria-current": "page", "href": "/"}).
			Append(nap.Text("Home")))
}

func InactiveNavItem(r nap.Router) nap.Elm {
	return r.
		E("li", nap.M{"class": "nav-item"}).
		Append(r.
			E("a", nap.M{"class": "nav-link", "href": "/"}).
			Append(nap.Text("Link")))
}

func DropdownNavItem(r nap.Router) nap.Elm {
	return r.
		E("li", nap.M{"class": "nav-item dropdown"}).
		Append(r.
			E("a", nap.M{
				"class":          "nav-link dropdown-toggle",
				"role":           "button",
				"data-bs-toggle": "dropdown",
				"aria-expanded":  "false",
			}).
			Append(nap.Text("Dropdown")), r.
			Elm(NavbarDropdownMenu))
}

func DropdownMenu(r nap.Router) nap.Elm {
	return r.
		E("ul", nap.M{"class": "dropdown-menu"}).
		Append(r.
			E("li").
			Append(r.
				E("a", nap.M{"class": "dropdown-item", "href": "/"}).
				Append(nap.Text("Action"))), r.
			E("li").
			Append(r.
				E("a", nap.M{"class": "dropdown-item", "href": "/"}).
				Append(nap.Text("Another action"))), r.
			E("li").
			Append(r.
				E("hr", nap.M{"class": "dropdown-divider"})), r.
			E("li").
			Append(r.
				E("a", nap.M{"class": "dropdown-item", "href": "/"}).
				Append(nap.Text("Something else here"))))
}

func DisabledNavItem(r nap.Router) nap.Elm {
	return r.
		E("li", nap.M{"class": "nav-item"}).
		Append(r.
			E("a", nap.M{"class": "nav-link disabled", "aria-disabled": "true"}).
			Append(nap.Text("Disabled")))
}

func SearchForm(r nap.Router) nap.Elm {
	return r.
		E("form", nap.M{"id": "searchForm", "class": "d-flex", "role": "search"}).
		Append(r.
			E("input", nap.M{
				"class":       "form-control me-2",
				"type":        "search",
				"placeholder": "Search",
				"aria-label":  "Search",
			}), r.
			E("button", nap.M{"class": "btn btn-outline-success", "type": "submit"}).
			Append(nap.Text("Search")))
}
