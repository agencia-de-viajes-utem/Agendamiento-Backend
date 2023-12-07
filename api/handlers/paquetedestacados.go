package handlers

import (
	"backend/api/models"
	"backend/api/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func ObtenerPaquetesDestacados(w http.ResponseWriter, r *http.Request) {
	db, err := utils.OpenDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
	WITH ranked_dates AS (
		SELECT
			fechapaquete.id,
			paquete.nombre,
			COALESCE(total_personas, 0) AS total_personas,
			fechapaquete.fechainit,
			fechapaquete.fechafin,
			ciudad_origen.id AS id_ciudad_origen,
			ciudad_destino.id AS id_ciudad_destino,
			ciudad_origen.nombre AS nombre_ciudad_origen,
			ciudad_destino.nombre AS nombre_ciudad_destino,
			fechapaquete.precio_oferta_vuelo as oferta_vuelo,
			paquete.precio_vuelo,
			habitacionhotel.precio_noche,
			paquete.descripcion,
			paquete.detalles,
			paquete.imagenes,
			opcionhotel.nombre AS nombre_opcion_hotel,
			habitacionhotel.descripcion AS descripcion_habitacion,
			habitacionhotel.servicios AS servicios_habitacion,
			hotel.nombre AS nombre_hotel,
			hotel.direccion AS direccion_hotel,
			hotel.valoracion AS valoracion_hotel,
			hotel.descripcion AS descripcion_hotel,
			hotel.servicios AS servicios_hotel,
			hotel.telefono AS telefono_hotel,
			hotel.correo_electronico AS correo_electronico_hotel,
			hotel.sitio_web AS sitio_web_hotel,
			-- aerolinea.nombre AS nombre_aerolinea, 
			-- aerolinea.imagen AS imagen_aerolinea,
			ROW_NUMBER() OVER (PARTITION BY fechapaquete.id ORDER BY fechapaquete.id) AS row_num
		FROM
			paquete
			INNER JOIN unnest(paquete.id_hh) WITH ORDINALITY t(habitacion_id, ord) ON TRUE
			INNER JOIN habitacionhotel ON t.habitacion_id = habitacionhotel.id
			INNER JOIN hotel ON habitacionhotel.hotel_id = hotel.id
			INNER JOIN opcionhotel ON habitacionhotel.opcion_hotel_id = opcionhotel.id
			INNER JOIN ciudad ciudad_origen ON paquete.id_origen = ciudad_origen.id
			INNER JOIN ciudad ciudad_destino ON paquete.id_destino = ciudad_destino.id
			INNER JOIN fechapaquete ON paquete.id = fechapaquete.id_paquete
			LEFT JOIN (
				SELECT
					paquete.id AS paquete_id,
					SUM(opcionhotel.cantidad) AS total_personas
				FROM
					paquete
					INNER JOIN unnest(paquete.id_hh) WITH ORDINALITY t(habitacion_id, ord) ON TRUE
					INNER JOIN habitacionhotel ON t.habitacion_id = habitacionhotel.id
					INNER JOIN opcionhotel ON habitacionhotel.opcion_hotel_id = opcionhotel.id
				GROUP BY
					paquete.id
			) AS subquery ON paquete.id = subquery.paquete_id
			LEFT JOIN aerolinea ON paquete.id = aerolinea.id_paquete  
	)
	SELECT *
	FROM ranked_dates
	WHERE 
    row_num = 1
		AND valoracion_hotel > 4.5
	ORDER BY valoracion_hotel DESC
	LIMIT 8
	`)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var paqueteDestacados []models.PaquetesDestacados

	for rows.Next() {
		var paqueteDestacado models.PaquetesDestacados
		var infoPaqueteDestacado models.PaqueteInfoAdicionalDestacado
		var hotelInfoDestacado models.HotelInfoDestacado

		err := rows.Scan(
			&paqueteDestacado.ID,
			&paqueteDestacado.Nombre,
			&paqueteDestacado.TotalPersonas,
			&paqueteDestacado.FechaInit,
			&paqueteDestacado.FechaFin,
			&paqueteDestacado.IdOrigen,
			&paqueteDestacado.IdDestino,
			&paqueteDestacado.NombreCiudadOrigen,
			&paqueteDestacado.NombreCiudadDestino,
			&paqueteDestacado.PrecioOfertaVuelo,
			&paqueteDestacado.PrecioVuelo,
			&paqueteDestacado.PrecioNoche,
			&paqueteDestacado.Descripcion,
			&paqueteDestacado.Detalles,
			&paqueteDestacado.Imagenes,
			&infoPaqueteDestacado.NombreOpcionHotel,
			&infoPaqueteDestacado.DescripcionHabitacion,
			&infoPaqueteDestacado.ServiciosHabitacion,
			&hotelInfoDestacado.NombreHotel,
			&hotelInfoDestacado.DireccionHotel,
			&hotelInfoDestacado.ValoracionHotel,
			&hotelInfoDestacado.DescripcionHotel,
			&hotelInfoDestacado.ServiciosHotel,
			&hotelInfoDestacado.TelefonoHotel,
			&hotelInfoDestacado.CorreoElectronico,
			&hotelInfoDestacado.SitioWeb,
			//&paqueteDestacado.Aerolinea,
			//&paqueteDestacado.AerolineaImagen,
			&infoPaqueteDestacado.RowNum,
		)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Error al escanear los resultados", http.StatusInternalServerError)
			return
		}
		fechaInicio, err := time.Parse("2006-01-02T15:04:05Z", paqueteDestacado.FechaInit)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Error al parsear la fecha de inicio", http.StatusInternalServerError)
			return
		}

		fechaFin, err := time.Parse("2006-01-02T15:04:05Z", paqueteDestacado.FechaFin)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Error al parsear la fecha de fin", http.StatusInternalServerError)
			return
		}

		paqueteDestacado.FechaInit = fechaInicio.Format("2006-01-02")
		paqueteDestacado.FechaFin = fechaFin.Format("2006-01-02")

		infoPaqueteDestacado.HotelInfo = hotelInfoDestacado
		paqueteDestacado.InfoPaquete = infoPaqueteDestacado

		paqueteDestacados = append(paqueteDestacados, paqueteDestacado)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(paqueteDestacados); err != nil {
		log.Fatal(err)
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
	}
}
