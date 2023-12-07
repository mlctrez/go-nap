package nav

import (
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/go-nap/nap/jsa"
	"strings"
)

func NavbarOverride(r nap.Router) {
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
}
