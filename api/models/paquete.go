package models

type PaqueteInfo struct {
	ID                  int     `json:"id"`
	Nombre              string  `json:"nombre"`
	TotalPersonas       int     `json:"total_personas"`
	FechaInit           string  `json:"fechainit"`
	FechaFin            string  `json:"fechafin"`
	IdOrigen            int     `json:"id_ciudad_origen"`
	IdDestino           int     `json:"id_ciudad_destino"`
	NombreCiudadOrigen  string  `json:"nombre_ciudad_origen"`
	NombreCiudadDestino string  `json:"nombre_ciudad_destino"`
	PrecioOfertaVuelo   float64 `json:"oferta_vuelo"`
	PrecioVuelo         float64 `json:"precio_vuelo"`
	PrecioNoche         float64 `json:"precio_noche"`
	Descripcion         string  `json:"descripcion"`
	Detalles            string  `json:"detalles"`
	Imagenes            string  `json:"imagenes"`
	//Aerolinea           string               `json:"aerolinea"`
	//AerolineaImagen     string               `json:"imagen_aerolinea"`
	InfoPaquete PaqueteInfoAdicional `json:"info_paquete"`
}

type PaqueteInfoAdicional struct {
	HabitacionId          int       `json:"habitacion_id"`
	OpcionHotelId         int       `json:"opcion_hotel_id"`
	NombreOpcionHotel     string    `json:"nombre_opcion_hotel"`
	DescripcionHabitacion string    `json:"descripcion_habitacion"`
	ServiciosHabitacion   string    `json:"servicios_habitacion"`
	HotelInfo             HotelInfo `json:"hotel_info"`
	RowNum                int       `json:"row_num"`
}

type HotelInfo struct {
	ID                int     `json:"id_hotel"`
	NombreHotel       string  `json:"nombre_hotel"`
	CiudadIdHotel     int     `json:"ciudad_id_hotel"`
	DireccionHotel    string  `json:"direccion_hotel"`
	ValoracionHotel   float64 `json:"valoracion_hotel"`
	DescripcionHotel  string  `json:"descripcion_hotel"`
	ServiciosHotel    string  `json:"servicios_hotel"`
	TelefonoHotel     string  `json:"telefono_hotel"`
	CorreoElectronico string  `json:"correo_electronico_hotel"`
	SitioWeb          string  `json:"sitio_web_hotel"`
}
