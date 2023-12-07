// middleware/cors.go
package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// CORSHandler crea un manejador CORS con las opciones dadas.
func CORSHandler(allowedOrigins []string, router *mux.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type"},
	})

	return c.Handler(router)
}

// RestrictedHandler devuelve un manejador que restringe el acceso a las rutas indicadas para métodos no GET.
func RestrictedHandler(restrictedRoutes map[string]bool, allowedOrigins []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verificar el origen de la solicitud
		origin := r.Header.Get("Origin")
		originAllowed := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				originAllowed = true
				break
			}
		}

		// Verificar si la ruta está permitida para métodos no GET
		if r.Method != http.MethodGet {
			if !originAllowed || !restrictedRoutes[r.URL.Path] {
				http.Error(w, "Not Allowed", http.StatusForbidden)
				return
			}
		}

		// Si el origen es permitido y la ruta es pública o permitida, pasar al siguiente manejador
		next.ServeHTTP(w, r)
	})
}
