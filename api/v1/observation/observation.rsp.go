package observation

type ObservationRsp struct {
	ID          uint   `json:"observation_id"`
	Nom         string `json:"nom"`
	Observation string `json:"observation"`
	Revue       string `json:"revue"`
	//Entretien             string `json:"entretien"`
	//VerificationTechnique string `json:"verification_technique"`
	//Analyse               string `json:"analyse"`
	ListeControleID uint `json:"listecontrole_id"`
}
