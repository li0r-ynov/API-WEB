package routes

import (
	"net/http"

	"pokemon/controllers"
)

// characterRoutes enregistre les routes de la partie personnages.
func characterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /pokemon", controllers.ListCharacters)
	mux.HandleFunc("GET /pokemon/{id}", controllers.ShowCharacter)
}
