package handlers

import (
	"backend/api/utils"
	"encoding/json"
	"net/http"
)

type SolicitudVista struct {
	Fk_fechaPaquete int `json:"fk_fechaPaquete"`
}

type Respuesta struct {
	Mensaje string `json:"mensaje"`
}

func AgregarVista(w http.ResponseWriter, r *http.Request) {
	var solicitudVista SolicitudVista
	err := json.NewDecoder(r.Body).Decode(&solicitudVista)
	if err != nil {
		http.Error(w, "Error al parsear la solicitud JSON", http.StatusBadRequest)
		return
	}

	db, err := utils.OpenDB()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
	INSERT INTO 
	logs_paquetes (fk_fechapaquete, cantidad_vistas)
 	VALUES ($1, 1)
	`, solicitudVista.Fk_fechaPaquete)

	if err != nil {
		http.Error(w, "Error al agregar la vista", http.StatusInternalServerError)
		return
	}

	// Send a JSON response indicating success
	respuesta := Respuesta{Mensaje: "Vista agregada correctamente"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
