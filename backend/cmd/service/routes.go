package main

func (app *App) setRoutes() {
	// app.router.HandleFunc("/ws", app.controller.WebsocketHandler)

	apis := app.router.PathPrefix("/api/").Subrouter()

	authApis := apis.PathPrefix("/auth/").Subrouter()
	// TODO: add not logged in middleware
	authApis.HandleFunc("/register", app.controller.AuthController.Register).Methods("POST")
}
