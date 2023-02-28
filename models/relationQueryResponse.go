package models

// RelationQueryResponse: tuebe el true o false que se obtiene de consultar la relacion entre dos usuarios
type RelationQueryResponse struct {
	Status bool `bson:"status" json:"status"`
}
