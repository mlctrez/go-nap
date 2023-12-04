package nav

import (
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/go-nap/nap/jsa"
)

func Override(r nap.Router, path string, el nap.Elm) nap.Elm {
	switch path {
	case "navbar/form":
		el.Listen("submit", jsa.FuncOf(func(this jsa.Value, args []jsa.Value) any {
			args[0].PreventDefault()
			input := el.First("input")
			if input.Get("value").String() != "" {
				jsa.Log("search value", input.Get("value"))
				input.Set("value", "")
			}
			return nil
		}))
		return el
	case "navbar/brand":
		return nap.El("a",
			nap.M{"class": "navbar-brand", "href": "https://github.com/mlctrez/go-nap"},
		).Append(nap.Text("go-nap"))
	}
	return el
}
