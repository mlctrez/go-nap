package compo

import (
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/go-nap/nap/jsa"
	"strings"
)

func NavbarOverride(r nap.Router) {
	r.Override(CNavbarNav, func(r nap.Router) nap.Elm {
		return r.NavLink(r.ElmOrig(CNavbarNav), "/", "NAP")
	})
	r.Override(CNavbarSearchForm, func(r nap.Router) nap.Elm {
		searchForm := r.ElmOrig(CNavbarSearchForm)
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
	r.Override(CNavbarItems, func(r nap.Router) nap.Elm {
		navItems := r.ElmOrig(CNavbarItems)
		path := jsa.CurrentURL().Path
		for _, ni := range navList {
			var item nap.Elm
			if ni.url == path {
				item = r.ElmOrig(CNavbarActiveNavItem)
			} else {
				item = r.ElmOrig(CNavbarInactiveNavItem)
			}
			navItems.Append(r.NavLink(item, ni.url, ni.name))
		}

		dropDown := r.Elm(CNavbarDropdownNavItem)
		ul := dropDown.First("ul")
		ul.Append(r.NavLink(r.Elm(CNavbarDropdownItem), "/", "home"))
		ul.Append(r.Elm(CNavbarDropdownDivider))
		ul.Append(r.NavLink(r.Elm(CNavbarDropdownItem), "/other", "other"))
		navItems.Append(dropDown)

		return navItems
	})
}
