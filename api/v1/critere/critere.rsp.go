package critere

type CritereRsp struct {
	Nom         string `json:"nom"`
	Description string `json:"description"`
	//Observation           string `json:"observation"`
	Entretien             string `json:"entretien"`
	VerificationTechnique string `json:"verification_technique"`
	Analyse               string `json:"analyse"`
	NormeID               uint   `json:"norme_id"`
}
