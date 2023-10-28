package models

// PublicationJob: captura el body y el mensaje con los campos respectivos del anuncio del empleo
type PublicationJob struct {
	Position        string `bson:"position" json:"position"`
	Company         string `bson:"company" json:"company"`
	TypeOfWorkplace string `bson:"typeOfWorkplace" json:"typeOfWorkplace"`
	JobLocation     string `bson:"jobLocation" json:"jobLocation"`
	JobType         string `bson:"jobType" json:"jobType"`
	Description     string `bson:"description" json:"description"`
	EmailRecruiter  string `bson:"emailRecruiter" json:"emailRecruiter"`
	Finished        bool   `bson:"finished" json:"finished"`
}
