package storeweb

type StoreResponse struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	NoHp    string `json:"number_handphone"`
	Address string `json:"address"`
	Balance int64  `json:"balance"`
}

type StoreResponseGlobal struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	NoHp    string `json:"number_handphone"`
	Address string `json:"address"`
}
