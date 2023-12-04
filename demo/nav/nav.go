package nav

import nap "github.com/mlctrez/go-nap/nap"

const NavbarNav = "navbar/nav"
const NavbarBrand = "navbar/brand"
const NavbarItems = "navbar/items"
const NavbarDropdown = "navbar/dropdown"
const NavbarDropdownMenu = "navbar/dropdownMenu"
const NavbarSearchForm = "navbar/searchForm"

func Navbar(r nap.Router) {
	r.ElmFunc(NavbarNav, Nav)
	r.ElmFunc(NavbarBrand, Brand)
	r.ElmFunc(NavbarItems, Items)
	r.ElmFunc(NavbarDropdown, Dropdown)
	r.ElmFunc(NavbarDropdownMenu, DropdownMenu)
	r.ElmFunc(NavbarSearchForm, SearchForm)
}

func Nav(r nap.Router) nap.Elm {
	return r.
		E("nav", nap.M{"class": "navbar navbar-expand-lg bg-body-tertiary", "data-nap-prefix": "navbar", "data-nap-key": "navbar"}).
		Append(r.
			E("div", nap.M{"class": "container-fluid"}).
			Append(r.
				Elm(NavbarBrand), r.
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
					Elm(NavbarItems), r.
					Elm(NavbarSearchForm))))
}

func Brand(r nap.Router) nap.Elm {
	return r.
		E("a", nap.M{"class": "navbar-brand", "href": "/"}).
		Append(nap.Text("Navbar foo"))
}

func Items(r nap.Router) nap.Elm {
	return r.
		E("ul", nap.M{"class": "navbar-nav me-auto mb-2 mb-lg-0"}).
		Append(r.
			E("li", nap.M{"class": "nav-item"}).
			Append(r.
				E("a", nap.M{"class": "nav-link active", "aria-current": "page", "href": "/"}).
				Append(nap.Text("Home"))), r.
			E("li", nap.M{"class": "nav-item"}).
			Append(r.
				E("a", nap.M{"class": "nav-link", "href": "/"}).
				Append(nap.Text("Link"))), r.
			Elm(NavbarDropdown), r.
			E("li", nap.M{"class": "nav-item"}).
			Append(r.
				E("a", nap.M{"class": "nav-link disabled", "aria-disabled": "true"}).
				Append(nap.Text("Disabled"))))
}

func Dropdown(r nap.Router) nap.Elm {
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
