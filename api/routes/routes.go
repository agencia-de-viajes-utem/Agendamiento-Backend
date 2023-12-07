package routes

import (
	"backend/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes configura las rutas en el enrutador proporcionado
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola Mundo!"))
	}).Methods(http.MethodGet)

	// Rutas para usuarios y mensajes
	// router.HandleFunc("/usuarios", handlers.GetAllUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/aeropuertos", handlers.ListarAeropuertos).Methods(http.MethodGet)
	router.HandleFunc("/paquetes", handlers.ObtenerPaquetes).Methods(http.MethodPost)        // pide un JSON en el body es POST
	router.HandleFunc("/paquetes/mes", handlers.ObtenerPaquetesMes).Methods(http.MethodPost) // pide un JSON en el body es POST
	router.HandleFunc("/paquetes/destacados", handlers.ObtenerPaquetesDestacados).Methods(http.MethodGet)
	router.HandleFunc("/paquetes/ofertas", handlers.ObtenerPaquetesOfertas).Methods(http.MethodPost) // pide un JSON en el body es POST
	router.HandleFunc("/paquetes/mas-vistos", handlers.ObtenerMasVistos).Methods(http.MethodGet)
	router.HandleFunc("/anadir", handlers.AgregarVista).Methods(http.MethodPost) // pide un JSON en el body es POST
	router.HandleFunc("/aeropuerto", handlers.ObtenerAeropuertos).Methods(http.MethodPost)
	router.HandleFunc("/comentario", handlers.ObtenerUsuarioYValoracion).Methods(http.MethodGet)

}
