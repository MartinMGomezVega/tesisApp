package models

// Relation: modelo para grabar la relacion de un usuario con otro
type Relation struct {
	UserID         string `bson:"userId" json:"userId"`
	UserRelationID string `bson:"userRelationId" json:"userRelationId"`
}
