package entreprise

type EntrepriseRsp struct {
	Nom       string `json:"nom"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	ManagedBy uint   `json:"managed_by"`
}
