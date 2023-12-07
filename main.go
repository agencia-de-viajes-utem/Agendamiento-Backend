// main.go
package main

import (
	"backend/api/middleware"
	"backend/api/routes"
	"backend/api/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Cargar las variables de entorno
	utils.LoadEnv()

	// Obtener el valor de la variable PORT o usar un valor predeterminado
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}

	// Crear el router
	r := mux.NewRouter()

	routes.RegisterRoutes(r)

	// Configurar CORS con el paquete 'middleware'
	handler := middleware.CORSHandler([]string{"http://localhost:3000"}, r)

	// Configurar un handler adicional para restringir a los orígenes permitidos y rutas restringidas
	restrictedRoutes := map[string]bool{
		"/paquetes":         true,
		"/paquetes/mes":     true,
		"/paquetes/ofertas": true,
		"/anadir":           true,
	}
	restrictedHandler := middleware.RestrictedHandler(restrictedRoutes, []string{"http://localhost:3000"}, handler)

	fmt.Printf("Servidor corriendo en http://localhost%s\n", port)
	if err := http.ListenAndServe(port, restrictedHandler); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
