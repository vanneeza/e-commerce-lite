package storeweb

type StoreCreateRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	NoHp    string `json:"number_handphone"`
	Address string `json:"address"`
}
