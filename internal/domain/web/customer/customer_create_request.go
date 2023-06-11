package customerweb

type CustomerCreateRequest struct {
	Name     string `json:"name"`
	NoHp     string `json:"number_handphone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}
