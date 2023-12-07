package nav

import (
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/go-nap/nap/jsa"
	"strings"
)

func NavbarOverride(r nap.Router) {
	r.Override(NavbarNav, func(r nap.Router) nap.Elm {
		return r.NavLink(r.ElmOrig(NavbarNav), "/", "NAP")
	})
	r.Override(NavbarSearchForm, func(r nap.Router) nap.Elm {
		searchForm := r.ElmOrig(NavbarSearchForm)
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
	r.Override(NavbarItems, func(r nap.Router) nap.Elm {
		navList := []struct {
			url  string
			name string
		}{
			{url: "/", name: "home"},
			{url: "/other", name: "other"},
		}
		navItems := r.ElmOrig(NavbarItems)
		path := jsa.CurrentURL().Path
		for _, ni := range navList {
			var item nap.Elm
			if ni.url == path {
				item = r.ElmOrig(NavbarActiveNavItem)
			} else {
				item = r.ElmOrig(NavbarInactiveNavItem)
			}
			navItems.Append(r.NavLink(item, ni.url, ni.name))
		}
		return navItems
	})
}
