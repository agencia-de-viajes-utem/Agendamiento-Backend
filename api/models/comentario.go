package models

type ValoracionInfo struct {
	Estrellas              int    `json:"estrellas"`
	Comentario             string `json:"comentario"`
	NombreUsuario          string `json:"nombre_usuario"`
	ApellidoUsuario        string `json:"apellido_usuario"`
	SegundoApellidoUsuario string `json:"segundoapellido_usuario"`
	ImgProfileUsuario      string `json:"img_profile_usuario"`
}
