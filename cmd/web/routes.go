package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/pharsha1995/snippetbox/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /ping", ping)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)
	protected := dynamic.Append(app.requireAuthentication)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /about", dynamic.ThenFunc(app.about))

	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))

	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))

	mux.Handle("GET /account/view", protected.ThenFunc(app.accountView))
	mux.Handle("GET /account/password/update", protected.ThenFunc(app.accountPasswordUpdate))
	mux.Handle("POST /account/password/update", protected.ThenFunc(app.accountPasswordUpdatePost))

	return standard.Then(mux)
}
