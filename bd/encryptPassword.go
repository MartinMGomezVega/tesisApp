package bd

import "golang.org/x/crypto/bcrypt"

// EncryptPassword: encriptar contraseña
func EncryptPassword(pass string) (string, error) {
	// costo: es la cantidad de pasadas para encriptar, a mayor costo mas encriptado pero mas demora. Y es el valor elevado.
	costo := 8                                                     // 8 serian 256 veces de pasadas
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo) // Procesa la password y devuelve un SLICE de bytes (SLICE es un vector sin tope)
	return string(bytes), err
}
