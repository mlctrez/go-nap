package web

import nap "github.com/mlctrez/go-nap/nap"

const CBaseHtml = "base/html"
const CBaseHeader = "base/header"

func Base(r nap.Router) {
	r.ElmFunc(CBaseHtml, BaseHtml)
	r.ElmFunc(CBaseHeader, BaseHeader)
	BaseOverride(r)
}

func BaseHtml(r nap.Router) nap.Elm {
	return r.E("html").Set("data-bs-theme", "dark").Set("lang", "en").
		Append(r.Elm(CBaseHeader))
}

func BaseHeader(r nap.Router) nap.Elm {
	return r.E("head").
		Append(
			r.E("meta").Set("content", "light dark").Set("name", "color-scheme"),
			r.E("link").Set("href", "logo.svg").Set("rel", "icon"),
			r.E("link").Set("href", "bootstrap.min.css").Set("rel", "stylesheet"),
			r.E("link").Set("href", "sign-in.css").Set("rel", "stylesheet"),
			r.E("script").Set("src", "bootstrap.bundle.min.js"),
			r.E("script").Set("src", "wasm.js"),
			r.E("script").Set("src", "runtime.js"),
			r.E("title").
				Append(nap.Text("demo")))
}
