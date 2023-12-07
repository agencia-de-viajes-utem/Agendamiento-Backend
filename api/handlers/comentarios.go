package handlers

import (
	"backend/api/models"
	"backend/api/utils"
	"encoding/json"
	"net/http"
)

// ObtenerUsuarioYValoracion is the handler for the GET request to retrieve user and rating information
func ObtenerUsuarioYValoracion(w http.ResponseWriter, r *http.Request) {
	id_paquete := r.URL.Query().Get("id_paquete")

	// Open the connection to the database
	db, err := utils.OpenDB()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT
		v.estrellas,
		v.comentario,
		u.nombre AS nombre_usuario,
		u.apellido AS apellido_usuario,
		u.segundoapellido AS segundoapellido_usuario,
		u.img_profile AS img_profile_usuario
	FROM
		valoracion v
	JOIN
		reserva r ON v.id_reserva = r.id
	JOIN
		usuario u ON v.id_usuario = u.id
	JOIN
		fechapaquete fp ON r.id_fechapaquete = fp.id
	JOIN
		paquete p ON fp.id_paquete = p.id
	WHERE
			p.id = $1`, id_paquete)
	if err != nil {
		http.Error(w, "Error al realizar la consulta", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var valoraciones []models.ValoracionInfo
	for rows.Next() {
		var valoracion models.ValoracionInfo
		if err := rows.Scan(
			&valoracion.Estrellas,
			&valoracion.Comentario,
			&valoracion.NombreUsuario,
			&valoracion.ApellidoUsuario,
			&valoracion.SegundoApellidoUsuario,
			&valoracion.ImgProfileUsuario,
		); err != nil {
			http.Error(w, "Error al escanear resultados", http.StatusInternalServerError)
			return
		}

		valoraciones = append(valoraciones, valoracion)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(valoraciones); err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
	}
}
