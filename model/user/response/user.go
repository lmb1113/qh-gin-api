package response

type LoginResponse struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}
