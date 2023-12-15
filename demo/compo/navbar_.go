package compo

import (
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/go-nap/nap/jsa"
	"strings"
)

func NavbarOverride(r nap.Router) {
	r.Override(ENavbarNav, func(r nap.Router) nap.Elm {
		return r.NavLink(r.ElmOrig(ENavbarNav), "/", "NAP")
	})
	r.Override(ENavbarSearchForm, func(r nap.Router) nap.Elm {
		searchForm := r.ElmOrig(ENavbarSearchForm)
		return searchForm.Listen("submit", jsa.FuncOf(func(this jsa.Value, args []jsa.Value) any {
			args[0].PreventDefault()
			searchValue := searchForm.First("input").Get("value")
			if strings.TrimSpace(searchValue.String()) != "" {
				jsa.Log(searchValue)
				searchForm.First("input").Set("value", "")
			}
			return nil
		}))
	})
	navList := []struct {
		url  string
		name string
	}{
		{url: "/", name: "home"},
		{url: "/other", name: "other"},
	}
	r.Override(ENavbarItems, func(r nap.Router) nap.Elm {
		navItems := r.ElmOrig(ENavbarItems)
		path := jsa.CurrentURL().Path
		for _, ni := range navList {
			var item nap.Elm
			if ni.url == path {
				item = r.ElmOrig(ENavbarActiveNavItem)
			} else {
				item = r.ElmOrig(ENavbarInactiveNavItem)
			}
			navItems.Append(r.NavLink(item, ni.url, ni.name))
		}

		dropDown := r.Elm(ENavbarDropdownNavItem)
		ul := dropDown.First("ul")
		ul.Append(r.NavLink(r.Elm(ENavbarDropdownItem), "/", "home"))
		ul.Append(r.Elm(ENavbarDropdownDivider))
		ul.Append(r.NavLink(r.Elm(ENavbarDropdownItem), "/other", "other"))
		navItems.Append(dropDown)

		return navItems
	})
}
