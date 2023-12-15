package compo

import nap "github.com/mlctrez/go-nap/nap"

const CSignInMain = "signIn/main"

func SignIn(r nap.Router) {
	r.ElmFunc(CSignInMain, SignInMain)
	SignInOverride(r)
}

func SignInMain(r nap.Router) nap.Elm {
	return r.E("main").Set("class", "form-signin w-100 m-auto").
		Append(r.E("form").
			Append(
				r.E("h1").Set("class", "h3 mb-3 fw-normal").
					Append(nap.Text("Please sign in")),
				r.E("div").Set("class", "form-floating").
					Append(
						r.E("input").
							Set("type", "email").
							Set("class", "form-control").
							Set("id", "floatingInput").
							Set("placeholder", "name@example.com"),
						r.E("label").Set("for", "floatingInput").
							Append(nap.Text("Email address"))),
				r.E("div").Set("class", "form-floating").
					Append(
						r.E("input").
							Set("type", "password").
							Set("class", "form-control").
							Set("id", "floatingPassword").
							Set("placeholder", "Password"),
						r.E("label").Set("for", "floatingPassword").
							Append(nap.Text("Password"))),
				r.E("div").Set("class", "form-check text-start my-3").
					Append(
						r.E("input").
							Set("class", "form-check-input").
							Set("type", "checkbox").
							Set("value", "remember-me").
							Set("id", "flexCheckDefault"),
						r.E("label").Set("class", "form-check-label").Set("for", "flexCheckDefault").
							Append(nap.Text("Remember me"))),
				r.E("button").Set("class", "btn btn-primary w-100 py-2").Set("type", "submit").
					Append(nap.Text("Sign in")),
				r.E("p").Set("class", "mt-4 mb-3 text-body-secondary").
					Append(nap.Text("© 2023"))))
}
