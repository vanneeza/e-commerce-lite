package customerweb

type CustomerUpdateRequest struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	NoHp     string `json:"number_handphone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}
