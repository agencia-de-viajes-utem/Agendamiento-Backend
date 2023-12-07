package models

type Aeropuerto struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Ciudad string `json:"ciudad"`
	Pais   string `json:"pais"`
}
