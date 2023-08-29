package rapport

type RapportRsp struct {
	Nom                string `json:"nom"`
	Nombreofconformite uint   `json:"nombreofconformite"`
	Process            string `json:"process"`
	Nombreofarticle    uint   `json:"nombreofarticle"`
	Description        string `json:"description"`
	ObservationID      uint   `json:"observation_id"`
}
