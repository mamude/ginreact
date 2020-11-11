package requests

// CustomerRequest request
type CustomerRequest struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	SessionID string `json:"sessionId"`
}

// UserLoginRequest request
type UserLoginRequest struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	SessionID string `json:"sessionId"`
}
