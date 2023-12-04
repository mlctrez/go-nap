package web

import nap "github.com/mlctrez/go-nap/nap"

const BaseHtml = "base/html"
const BaseHeader = "base/header"
const BaseBody = "base/body"

func Base(r nap.Router) {
	r.ElmFunc(BaseHtml, Html)
	r.ElmFunc(BaseHeader, Header)
	r.ElmFunc(BaseBody, Body)
}

func Html(r nap.Router) nap.Elm {
	return r.
		E("html", nap.M{"data-bs-theme": "dark", "lang": "en", "data-nap-prefix": "base"}).
		Append(r.
			Elm(BaseHeader), r.
			Elm(BaseBody))
}

func Header(r nap.Router) nap.Elm {
	return r.
		E("head").
		Append(r.
			E("meta", nap.M{"content": "light dark", "name": "color-scheme"}), r.
			E("link", nap.M{"href": "logo.svg", "rel": "icon"}), r.
			E("link", nap.M{"href": "bootstrap.min.css", "rel": "stylesheet"}), r.
			E("script", nap.M{"src": "bootstrap.bundle.min.js"}), r.
			E("script", nap.M{"src": "wasm.js"}), r.
			E("script", nap.M{"src": "runtime.js"}), r.
			E("title").
			Append(nap.Text("demo")))
}

func Body(r nap.Router) nap.Elm {
	return r.
		E("body")
}
