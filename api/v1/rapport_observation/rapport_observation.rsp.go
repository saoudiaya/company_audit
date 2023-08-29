package rapportobservation

type RapportObservationRsp struct {
	Nom           string `json:"nom"`
	Description   string `json:"description"`
	ObservationID uint   `json:"observation_id"`
	RapportID     uint   `json:"rapport_id"`
}
