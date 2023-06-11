package loginweb

type LoginRequest struct {
	NoHp     string `json:"number_handphone"`
	Password string `json:"password"`
}
